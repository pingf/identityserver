package organization

import (
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"sort"

	"github.com/gorilla/context"
	"github.com/itsyouonline/identityserver/db"
	contractdb "github.com/itsyouonline/identityserver/db/contract"
	"github.com/itsyouonline/identityserver/db/organization"
	"github.com/itsyouonline/identityserver/db/user"
	validationdb "github.com/itsyouonline/identityserver/db/validation"
	"github.com/itsyouonline/identityserver/identityservice/contract"
	"github.com/itsyouonline/identityserver/identityservice/invitations"
	"github.com/itsyouonline/identityserver/oauthservice"
	"gopkg.in/mgo.v2"
)

const (
	itsyouonlineGlobalID                    = "itsyouonline"
	MAX_ORGANIZATIONS_PER_USER              = 1000
	MAX_AMOUNT_INVITATIONS_PER_ORGANIZATION = 10000
)

// OrganizationsAPI is the implementation for /organizations root endpoint
type OrganizationsAPI struct {
}

// byGlobalID implements sort.Interface for []Organization based on
// the GlobalID field.
type byGlobalID []organization.Organization

func (a byGlobalID) Len() int           { return len(a) }
func (a byGlobalID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byGlobalID) Less(i, j int) bool { return a[i].Globalid < a[j].Globalid }

// GetOrganizationTree is the handler for GET /organizations/{globalid}/tree
// Get organization tree.
func (api OrganizationsAPI) GetOrganizationTree(w http.ResponseWriter, r *http.Request) {
	var requestedOrganization = mux.Vars(r)["globalid"]
	//TODO: validate input
	parentGlobalID := ""
	var parentGlobalIDs = make([]string, 0, 1)
	for _, localParentID := range strings.Split(requestedOrganization, ".") {
		if parentGlobalID == "" {
			parentGlobalID = localParentID
		} else {
			parentGlobalID = parentGlobalID + "." + localParentID
		}

		parentGlobalIDs = append(parentGlobalIDs, parentGlobalID)
	}

	orgMgr := organization.NewManager(r)

	parentOrganizations, err := orgMgr.GetOrganizations(parentGlobalIDs)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	suborganizations, err := orgMgr.GetSubOrganizations(requestedOrganization)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	allOrganizations := append(parentOrganizations, suborganizations...)

	sort.Sort(byGlobalID(allOrganizations))

	//Build a treestructure
	var orgTree *OrganizationTreeItem
	orgTreeIndex := make(map[string]*OrganizationTreeItem)
	for _, org := range allOrganizations {
		newTreeItem := &OrganizationTreeItem{GlobalID: org.Globalid, Children: make([]*OrganizationTreeItem, 0, 0)}
		orgTreeIndex[org.Globalid] = newTreeItem
		if orgTree == nil {
			orgTree = newTreeItem
		} else {
			path := strings.Split(org.Globalid, ".")
			localName := path[len(path)-1]
			parentTreeItem := orgTreeIndex[strings.TrimSuffix(org.Globalid, "."+localName)]
			parentTreeItem.Children = append(parentTreeItem.Children, newTreeItem)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orgTree)
}

// CreateNewOrganization is the handler for POST /organizations
// Create a new organization. 1 user should be in the owners list. Validation is performed
// to check if the securityScheme allows management on this user.
func (api OrganizationsAPI) CreateNewOrganization(w http.ResponseWriter, r *http.Request) {
	var org organization.Organization

	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		log.Debug("Error decoding the organization:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if strings.Contains(org.Globalid, ".") {
		log.Debug("globalid contains a '.'")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	api.actualOrganizationCreation(org, w, r)

}

// CreateNewSubOrganization is the handler for POST /organizations/{globalid}/suborganizations
// Create a new suborganization.
func (api OrganizationsAPI) CreateNewSubOrganization(w http.ResponseWriter, r *http.Request) {
	parent := mux.Vars(r)["globalid"]
	var org organization.Organization

	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		log.Debug("Error decoding the organization:", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(org.Globalid, parent+".") {
		log.Debug("GlobalID does not start with the parent globalID")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	localid := strings.TrimPrefix(org.Globalid, parent+".")
	if strings.Contains(localid, ".") {
		log.Debug("localid contains a '.'")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	api.actualOrganizationCreation(org, w, r)

}

func (api OrganizationsAPI) actualOrganizationCreation(org organization.Organization, w http.ResponseWriter, r *http.Request) {

	if strings.TrimSpace(org.Globalid) == itsyouonlineGlobalID {
		log.Debug("Duplicate organization")
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if !org.IsValid() {
		log.Debug("Invalid organization")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	username := context.Get(r, "authenticateduser").(string)
	orgMgr := organization.NewManager(r)
	count, err := orgMgr.CountByUser(username)
	if err != nil {
		log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if count >= MAX_ORGANIZATIONS_PER_USER {
		log.Error("Reached organization limit for user ", username)
		writeErrorResponse(w, 422, "maximum_amount_of_organizations_reached")
		return
	}
	err = orgMgr.Create(&org)

	if err != nil && err != db.ErrDuplicate {
		log.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if err == db.ErrDuplicate {
		log.Debug("Duplicate organization")
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(&org)
}

// Get organization info
// It is handler for GET /organizations/{globalid}
func (api OrganizationsAPI) globalidGet(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]
	orgMgr := organization.NewManager(r)

	org, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}
	json.NewEncoder(w).Encode(org)
}

// Update organization info
// It is handler for PUT /organizations/{globalid}
func (api OrganizationsAPI) globalidPut(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]

	var org organization.Organization

	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	orgMgr := organization.NewManager(r)

	oldOrg, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}

	if org.Globalid != globalid {
		http.Error(w, "Changing globalid or id is Forbidden!", http.StatusForbidden)
		return
	}

	// Update only certain fields
	oldOrg.PublicKeys = org.PublicKeys
	oldOrg.DNS = org.DNS

	if err := orgMgr.Save(oldOrg); err != nil {
		log.Error("Error while saving organization: ", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(oldOrg)
}

// Assign a member to organization
// It is handler for POST /organizations/{globalid}/members
func (api OrganizationsAPI) globalidmembersPost(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]

	var s searchMember

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	orgMgr := organization.NewManager(r)

	org, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}

	// Check if user exists
	u, err := SearchUser(r, s.SearchString)
	if err != nil {
		if err != mgo.ErrNotFound {
			log.Error(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	for _, membername := range org.Members {
		if membername == u.Username {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	for _, membername := range org.Owners {
		if membername == u.Username {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}

	// Create JoinRequest
	invitationMgr := invitations.NewInvitationManager(r)
	count, err := invitationMgr.CountByOrganization(globalid)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if count >= MAX_AMOUNT_INVITATIONS_PER_ORGANIZATION {
		log.Error("Reached invitation limit for organization ", globalid)
		writeErrorResponse(w, 422, "max_amount_of_invitations_reached")
		return
	}
	orgReq := &invitations.JoinOrganizationInvitation{
		Role:         invitations.RoleMember,
		Organization: globalid,
		User:         u.Username,
		Status:       invitations.RequestPending,
	}

	if err := invitationMgr.Save(orgReq); err != nil {
		log.Error("Error inviting member: ", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(orgReq)
}

func (api OrganizationsAPI) UpdateOrganizationMemberShip(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]
	var membership Membership
	if err := json.NewDecoder(r.Body).Decode(&membership); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	orgMgr := organization.NewManager(r)
	org, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "updating organization membership", err)
		}
		return
	}
	var oldRole string
	for _, v := range org.Members {
		if v == membership.Username {
			oldRole = "members"
		}
	}
	for _, v := range org.Owners {
		if v == membership.Username {
			oldRole = "owners"
		}
	}
	err = orgMgr.UpdateMembership(globalid, membership.Username, oldRole, membership.Role)
	if err != nil {
		handleServerError(w, "updating organization membership", err)
		return
	}
	org, err = orgMgr.GetByName(globalid)
	if err != nil {
		handleServerError(w, "getting organization", err)
	}
	json.NewEncoder(w).Encode(org)

}

// RemoveOrganizationMember Remove a member from organization
// It is handler for DELETE /organizations/{globalid}/members/{username}
func (api OrganizationsAPI) RemoveOrganizationMember(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]
	username := mux.Vars(r)["username"]

	orgMgr := organization.NewManager(r)

	org, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}
	if err := orgMgr.RemoveMember(org, username); err != nil {
		log.Error("Error adding member: ", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// It is handler for POST /organizations/{globalid}/members
func (api OrganizationsAPI) globalidownersPost(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]

	var s searchMember

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	orgMgr := organization.NewManager(r)

	org, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}

	u, err := SearchUser(r, s.SearchString)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	for _, membername := range org.Owners {
		if membername == u.Username {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			return
		}
	}
	// Create JoinRequest
	invitationMgr := invitations.NewInvitationManager(r)

	orgReq := &invitations.JoinOrganizationInvitation{
		Role:         invitations.RoleOwner,
		Organization: globalid,
		User:         u.Username,
		Status:       invitations.RequestPending,
	}

	if err := invitationMgr.Save(orgReq); err != nil {
		log.Error("Error inviting owner: ", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(orgReq)
}

// RemoveOrganizationOwner Remove a member from organization
// It is handler for DELETE /organizations/{globalid}/owners/{username}
func (api OrganizationsAPI) RemoveOrganizationOwner(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]
	username := mux.Vars(r)["username"]

	orgMgr := organization.NewManager(r)

	org, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}

	if err := orgMgr.RemoveOwner(org, username); err != nil {
		log.Error("Error removing owner: ", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetPendingInvitations is the handler for GET /organizations/{globalid}/invitations
// Get the list of pending invitations for users to join this organization.
func (api OrganizationsAPI) GetPendingInvitations(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]

	invitationMgr := invitations.NewInvitationManager(r)

	requests, err := invitationMgr.GetPendingByOrganization(globalid)

	if err != nil {
		log.Error("Error in GetPendingByOrganization: ", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	pendingInvites := make([]organization.Invitation, len(requests), len(requests))
	for index, request := range requests {
		pendingInvites[index] = organization.Invitation{
			Role: request.Role,
			User: request.User,
		}
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(pendingInvites)
}

// RemovePendingInvitation is the handler for DELETE /organizations/{globalid}/invitations/{username}
// Cancel a pending invitation.
func (api OrganizationsAPI) RemovePendingInvitation(w http.ResponseWriter, r *http.Request) {
	log.Error("RemovePendingInvitation is not implemented")
}

// GetContracts is the handler for GET /organizations/{globalid}/contracts
// Get the contracts where the organization is 1 of the parties. Order descending by
// date.
func (api OrganizationsAPI) GetContracts(w http.ResponseWriter, r *http.Request) {
	globalID := mux.Vars(r)["globalId"]
	includedparty := contractdb.Party{Type: "org", Name: globalID}
	contract.FindContracts(w, r, includedparty)
}

// RegisterNewContract is handler for GET /organizations/{globalId}/contracts
func (api OrganizationsAPI) RegisterNewContract(w http.ResponseWriter, r *http.Request) {
	globalID := mux.Vars(r)["glabalId"]
	includedparty := contractdb.Party{Type: "org", Name: globalID}
	contract.CreateContract(w, r, includedparty)
}

// GetAPIKeyLabels is the handler for GET /organizations/{globalid}/apikeys
// Get the list of active api keys. The secrets themselves are not included.
func (api OrganizationsAPI) GetAPIKeyLabels(w http.ResponseWriter, r *http.Request) {
	organization := mux.Vars(r)["globalid"]

	mgr := oauthservice.NewManager(r)
	labels, err := mgr.GetClientLabels(organization)
	if err != nil {
		log.Error("Error getting a client secret labels: ", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(labels)
}

func isValidAPIKeyLabel(label string) (valid bool) {
	valid = true
	labelLength := len(label)
	valid = valid && labelLength > 1 && labelLength < 51
	return valid
}

func isValidDNSName(label string) (valid bool) {
	valid = true
	labelLength := len(label)
	valid = valid && labelLength > 2 && labelLength < 250
	return valid
}

// GetAPIKey is the handler for GET /organizations/{globalid}/apikeys/{label}
func (api OrganizationsAPI) GetAPIKey(w http.ResponseWriter, r *http.Request) {
	organization := mux.Vars(r)["globalid"]
	label := mux.Vars(r)["label"]

	mgr := oauthservice.NewManager(r)
	client, err := mgr.GetClient(organization, label)
	if err != nil {
		log.Error("Error getting a client: ", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if client == nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	apiKey := FromOAuthClient(client)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(apiKey)
}

// CreateNewAPIKey is the handler for POST /organizations/{globalid}/apikeys
// Create a new API Key, a secret itself should not be provided, it will be generated
// serverside.
func (api OrganizationsAPI) CreateNewAPIKey(w http.ResponseWriter, r *http.Request) {
	organization := mux.Vars(r)["globalid"]

	apiKey := APIKey{}

	if err := json.NewDecoder(r.Body).Decode(&apiKey); err != nil {
		log.Debug("Error decoding apikey: ", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	//TODO: validate key, not just the label property
	if !isValidAPIKeyLabel(apiKey.Label) {
		log.Debug("Invalid label: ", apiKey.Label)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	log.Debug("Creating apikey:", apiKey)
	c := oauthservice.NewOauth2Client(organization, apiKey.Label, apiKey.CallbackURL, apiKey.ClientCredentialsGrantType)

	mgr := oauthservice.NewManager(r)
	err := mgr.CreateClient(c)
	if db.IsDup(err) {
		log.Debug("Duplicate label")
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}
	if err != nil {
		log.Error("Error creating api secret label", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	apiKey.Secret = c.Secret

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(apiKey)

}

// UpdateAPIKey is the handler for PUT /organizations/{globalid}/apikeys/{label}
// Updates the label or other properties of a key.
func (api OrganizationsAPI) UpdateAPIKey(w http.ResponseWriter, r *http.Request) {
	organization := mux.Vars(r)["globalid"]
	oldlabel := mux.Vars(r)["label"]

	apiKey := APIKey{}

	if err := json.NewDecoder(r.Body).Decode(&apiKey); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if !isValidAPIKeyLabel(apiKey.Label) {
		log.Debug("Invalid label: ", apiKey.Label)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	mgr := oauthservice.NewManager(r)
	err := mgr.UpdateClient(organization, oldlabel, apiKey.Label, apiKey.CallbackURL, apiKey.ClientCredentialsGrantType)

	if err != nil && db.IsDup(err) {
		log.Debug("Duplicate label")
		http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
		return
	}

	if err != nil {
		log.Error("Error renaming api secret label", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteAPIKey is the handler for DELETE /organizations/{globalid}/apikeys/{label}
// Removes an API key
func (api OrganizationsAPI) DeleteAPIKey(w http.ResponseWriter, r *http.Request) {
	organization := mux.Vars(r)["globalid"]
	label := mux.Vars(r)["label"]

	mgr := oauthservice.NewManager(r)
	err := mgr.DeleteClient(organization, label)

	if err != nil {
		log.Error("Error deleting organization:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (api OrganizationsAPI) CreateDns(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]
	dnsName := mux.Vars(r)["dnsname"]

	if !isValidDNSName(dnsName) {
		log.Debug("Invalid DNS name: ", dnsName)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	orgMgr := organization.NewManager(r)
	organization, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}
	err = orgMgr.AddDNS(organization, dnsName)

	if err != nil {
		log.Error("Error creating DNS name", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := struct {
		Name string `json:"name"`
	}{
		Name: dnsName,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (api OrganizationsAPI) UpdateDns(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]
	oldDns := mux.Vars(r)["dnsname"]

	body := struct {
		Name string
	}{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if !isValidDNSName(body.Name) {
		log.Debug("Invalid DNS name: ", body.Name)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	orgMgr := organization.NewManager(r)
	organization, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}
	err = orgMgr.UpdateDNS(organization, oldDns, body.Name)

	if err != nil {
		log.Error("Error updating DNS name", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response := struct {
		Name string `json:"name"`
	}{
		Name: body.Name,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (api OrganizationsAPI) DeleteDns(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]
	dnsName := mux.Vars(r)["dnsname"]

	orgMgr := organization.NewManager(r)
	organization, err := orgMgr.GetByName(globalid)
	if err != nil {
		if err == mgo.ErrNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			handleServerError(w, "getting organization", err)
		}
		return
	}
	sort.Strings(organization.DNS)
	if sort.SearchStrings(organization.DNS, dnsName) == len(organization.DNS) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	err = orgMgr.RemoveDNS(organization, dnsName)

	if err != nil {
		log.Error("Error removing DNS name", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusNoContent)
}

// DeleteOrganization is the handler for DELETE /organizations/{globalid}
// Deletes an organization and all data linked to it (join-organization-invitations, oauth_access_tokens, oauth_clients)
func (api OrganizationsAPI) DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	globalid := mux.Vars(r)["globalid"]
	orgMgr := organization.NewManager(r)
	if !orgMgr.Exists(globalid) {
		writeErrorResponse(w, http.StatusNotFound, "organization_not_found")
		return
	}
	suborganizations, err := orgMgr.GetSubOrganizations(globalid)
	if handleServerError(w, "fetching suborganizations", err) {
		return
	}
	if len(suborganizations) != 0 {
		writeErrorResponse(w, 422, "organization_has_children")
		return
	}
	err = orgMgr.Remove(globalid)
	if handleServerError(w, "removing organization", err) {
		return
	}
	orgReqMgr := invitations.NewInvitationManager(r)
	err = orgReqMgr.RemoveAll(globalid)
	if handleServerError(w, "removing organization invitations", err) {
		return
	}

	oauthMgr := oauthservice.NewManager(r)
	err = oauthMgr.RemoveTokensByGlobalId(globalid)
	if handleServerError(w, "removing organization oauth accesstokens", err) {
		return
	}
	err = oauthMgr.RemoveClientsById(globalid)
	if handleServerError(w, "removing organization oauth clients", err) {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeErrorResponse(responseWriter http.ResponseWriter, httpStatusCode int, message string) {
	log.Debug(httpStatusCode, message)
	errorResponse := struct {
		Error string `json:"error"`
	}{message}
	responseWriter.WriteHeader(httpStatusCode)
	json.NewEncoder(responseWriter).Encode(&errorResponse)
}

func handleServerError(responseWriter http.ResponseWriter, actionText string, err error) bool {
	if err != nil {
		log.Error("Error while "+actionText, err)
		http.Error(responseWriter, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return true
	}
	return false
}

func SearchUser(r *http.Request, searchString string) (usr *user.User, err1 error) {
	userMgr := user.NewManager(r)
	usr, err1 = userMgr.GetByName(searchString)
	if err1 == mgo.ErrNotFound {
		valMgr := validationdb.NewManager(r)
		validatedPhonenumber, err2 := valMgr.GetByPhoneNumber(searchString)
		if err2 == mgo.ErrNotFound {
			validatedEmailAddress, err3 := valMgr.GetByEmailAddress(searchString)
			if err3 != nil {
				return nil, err3
			} else {
				return userMgr.GetByName(validatedEmailAddress.Username)
			}
		} else {
			return userMgr.GetByName(validatedPhonenumber.Username)
		}
	} else {
		return usr, err1
	}
}
