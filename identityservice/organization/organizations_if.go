package organization

//This file is auto-generated by go-raml
//Do not edit this file by hand since it will be overwritten during the next generation

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// OrganizationsInterface is interface for /organizations root endpoint
type OrganizationsInterface interface { // CreateNewOrganization is the handler for POST /organizations
	// Create a new organization. 1 user should be in the owners list. Validation is performed
	// to check if the securityScheme allows management on this user.
	CreateNewOrganization(http.ResponseWriter, *http.Request)
	// globalidGet is the handler for GET /organizations/{globalid}
	// Get organization info
	globalidGet(http.ResponseWriter, *http.Request)
	// CreateNewSubOrganization is the handler for POST /organizations/{globalid}/suborganizations
	// Create a new suborganization.
	CreateNewSubOrganization(http.ResponseWriter, *http.Request)
	// globalidPut is the handler for PUT /organizations/{globalid}
	// Update organization info
	globalidPut(http.ResponseWriter, *http.Request)
	// globalidmembersPost is the handler for POST /organizations/{globalid}/members
	// Assign a member to organization.
	globalidmembersPost(http.ResponseWriter, *http.Request)
	// globalidmembersusernameDelete is the handler for DELETE /organizations/{globalid}/members/{username}
	// Remove a member from organization
	globalidmembersusernameDelete(http.ResponseWriter, *http.Request)
	// globalidownersPost is the handler for POST /organizations/{globalid}/owners
	// Invite a user to become owner of an organization.
	globalidownersPost(http.ResponseWriter, *http.Request)
	// globalidownersusernameDelete is the handler for DELETE /organizations/{globalid}/owners/{username}
	// Remove an owner from organization
	globalidownersusernameDelete(http.ResponseWriter, *http.Request)
	// GetContracts is the handler for GET /organizations/{globalid}/contracts
	// Get the contracts where the organization is 1 of the parties. Order descending by
	// date.
	GetContracts(http.ResponseWriter, *http.Request)
	// GetPendingInvitations is the handler for GET /organizations/{globalid}/invitations
	// Get the list of pending invitations for users to join this organization.
	GetPendingInvitations(http.ResponseWriter, *http.Request)
	// RemovePendingInvitation is the handler for DELETE /organizations/{globalid}/invitations/{username}
	// Cancel a pending invitation.
	RemovePendingInvitation(http.ResponseWriter, *http.Request)
	// GetAPISecretLabels is the handler for GET /organizations/{globalid}/apisecrets
	// Get the list of labels that are defined for active api secrets. The secrets themselves
	// are not included.
	GetAPISecretLabels(http.ResponseWriter, *http.Request)
	// CreateNewAPISecret is the handler for POST /organizations/{globalid}/apisecrets
	// Create a new API Secret
	CreateNewAPISecret(http.ResponseWriter, *http.Request)
	// GetAPISecret is the handler for GET /organizations/{globalid}/apisecrets/{label}
	GetAPISecret(http.ResponseWriter, *http.Request)
	// UpdateAPISecretLabel is the handler for PUT /organizations/{globalid}/apisecrets/{label}
	// Updates the label of the secret
	UpdateAPISecretLabel(http.ResponseWriter, *http.Request)
	// DeleteAPISecret is the handler for DELETE /organizations/{globalid}/apisecrets/{label}
	// Removes an API secret
	DeleteAPISecret(http.ResponseWriter, *http.Request)
	// GetOrganizationTree is the handler for GET /organizations/{globalid}/tree
	GetOrganizationTree(http.ResponseWriter, *http.Request)
}

// OrganizationsInterfaceRoutes is routing for /organizations root endpoint
func OrganizationsInterfaceRoutes(r *mux.Router, i OrganizationsInterface) {
	r.Handle("/organizations", alice.New(newOauth2oauth_2_0Middleware([]string{}).Handler).Then(http.HandlerFunc(i.CreateNewOrganization))).Methods("POST")
	r.Handle("/organizations/{globalid}", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:member", "organization:owner"}).Handler).Then(http.HandlerFunc(i.globalidGet))).Methods("GET")
	r.Handle("/organizations/{globalid}", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.globalidPut))).Methods("PUT")
	r.Handle("/organizations/{globalid}/suborganizations", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.CreateNewSubOrganization))).Methods("POST")
	r.Handle("/organizations/{globalid}/members", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.globalidmembersPost))).Methods("POST")
	r.Handle("/organizations/{globalid}/members/{username}", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.globalidmembersusernameDelete))).Methods("DELETE")
	r.Handle("/organizations/{globalid}/owners", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.globalidownersPost))).Methods("POST")
	r.Handle("/organizations/{globalid}/owners/{username}", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.globalidownersusernameDelete))).Methods("DELETE")
	r.Handle("/organizations/{globalid}/contracts", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner", "organization:contracts:read"}).Handler).Then(http.HandlerFunc(i.GetContracts))).Methods("GET")
	r.Handle("/organizations/{globalid}/invitations", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.GetPendingInvitations))).Methods("GET")
	r.Handle("/organizations/{globalid}/invitations/{username}", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.RemovePendingInvitation))).Methods("DELETE")
	r.Handle("/organizations/{globalid}/apisecrets", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.GetAPISecretLabels))).Methods("GET")
	r.Handle("/organizations/{globalid}/apisecrets", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.CreateNewAPISecret))).Methods("POST")
	r.Handle("/organizations/{globalid}/apisecrets/{label}", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.GetAPISecret))).Methods("GET")
	r.Handle("/organizations/{globalid}/apisecrets/{label}", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.UpdateAPISecretLabel))).Methods("PUT")
	r.Handle("/organizations/{globalid}/apisecrets/{label}", alice.New(newOauth2oauth_2_0Middleware([]string{"organization:owner"}).Handler).Then(http.HandlerFunc(i.DeleteAPISecret))).Methods("DELETE")
	r.Handle("/organizations/{globalid}/tree", alice.New(newOauth2oauth_2_0Middleware([]string{}).Handler).Then(http.HandlerFunc(i.GetOrganizationTree))).Methods("GET")
}
