package user

import (
	"net/http"
	"strings"

	"crypto/ecdsa"
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/itsyouonline/identityserver/oauthservice"
)

// Oauth2oauth_2_0Middleware is oauth2 middleware for oauth_2_0
type Oauth2oauth_2_0Middleware struct {
	DescribedBy string
	Field       string
	Scopes      []string
}

//JWTPublicKey has the public key of the allowed JWT issuer
var JWTPublicKey ecdsa.PublicKey

// newOauth2oauth_2_0Middlewarecreate new Oauth2oauth_2_0Middleware struct
func newOauth2oauth_2_0Middleware(scopes []string) *Oauth2oauth_2_0Middleware {
	om := Oauth2oauth_2_0Middleware{
		Scopes:      scopes,
		DescribedBy: "headers",
		Field:       "Authorization",
	}
	return &om
}

// CheckScopes checks whether user has needed scopes
func (om *Oauth2oauth_2_0Middleware) CheckScopes(scopes []string) bool {
	if len(om.Scopes) == 0 {
		return true
	}

	for _, allowed := range om.Scopes {
		for _, scope := range scopes {
			if scope == allowed {
				return true
			}
		}
	}
	return false
}

// Handler return HTTP handler representation of this middleware
func (om *Oauth2oauth_2_0Middleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accessToken string
		var atscopestring string
		var username string
		var clientID string
		scopes := []string{}

		// access token checking
		if om.DescribedBy == "queryParameters" {
			accessToken = r.URL.Query().Get(om.Field)
		} else if om.DescribedBy == "headers" {
			accessToken = r.Header.Get(om.Field)
		}
		//Get the actual token out of the header (accept 'token ABCD' as well as just 'ABCD' and ignore some possible whitespace)
		if strings.HasPrefix(accessToken, "bearer") {
			jwtstring := strings.TrimSpace(strings.TrimPrefix(accessToken, "bearer"))
			token, err := jwt.Parse(jwtstring, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:

				if token.Method != jwt.SigningMethodES384 {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return &JWTPublicKey, nil
			})
			if err != nil || !token.Valid {
				log.Error(err)
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			username = token.Claims["username"].(string)
			clientID = token.Claims["aud"].(string)
			atscopestring = token.Claims["scope"].(string)

		} else if strings.HasPrefix(accessToken, "token") {
			accessToken = strings.TrimSpace(strings.TrimPrefix(accessToken, "token"))
			if accessToken == "" {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}

			log.Debug("Access Token: ", accessToken)
			//TODO: cache
			oauthMgr := oauthservice.NewManager(r)
			at, err := oauthMgr.GetAccessToken(accessToken)
			if err != nil {
				log.Error(err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			if at == nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			username = at.Username
			atscopestring = at.Scope
			clientID = at.ClientID
		} else {
			if webuser, ok := context.GetOk(r, "webuser"); ok {
				if parsedusername, ok := webuser.(string); ok && parsedusername != "" {
					username = parsedusername
					atscopestring = "admin"
					clientID = "itsyouonline"
				}
			}
		}
		if username == "" || clientID == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		protectedUsername := mux.Vars(r)["username"]

		if protectedUsername == username && clientID == "itsyouonline" && atscopestring == "admin" {
			scopes = append(scopes, "user:admin")
		}
		if strings.HasPrefix(atscopestring, "user:") {
			scopes = append(scopes, "user:info")
		}
		for _, scope := range strings.Split(atscopestring, ",") {
			scope = strings.Trim(scope, " ")
			scopes = append(scopes, scope)
		}

		log.Debug("Available scopes: ", scopes)

		context.Set(r, "client_id", clientID)
		context.Set(r, "availablescopes", atscopestring)

		// check scopes
		if !om.CheckScopes(scopes) {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
