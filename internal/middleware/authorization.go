package middleware

import (
	"errors"
	"github.com/renvins/gocoins_restful/internal/tools"
	"net/http"

	"github.com/renvins/gocoins_restful/api"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Response writer to write the response (body, headers, status code)
		// Request to get the request data (headers, body, method, URL)
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		var err error
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w, err)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}
