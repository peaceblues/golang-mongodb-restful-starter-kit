package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
)

// Headers set header to request
func Headers(r http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"})
	return handlers.CORS(headersOk, originsOk, methodsOk)(r)
}

// Response will return json responce of http
// This func handle both error a well as success
func Response(w http.ResponseWriter, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(payload)
}

func NewHttpError(errorCode string, statusCode int) map[string]interface{} {

	m := make(map[string]interface{})
	m["error"] = errorCode
	m["error_description"], _ = errorMessage[errorCode]
	m["code"] = statusCode

	return m
}

// Errors
const (
	InvalidUserID       = "invalidUserID"
	InternalError       = "internalError"
	UserNotFound        = "userNotFound"
	InvalidBindingModel = "invalidBindingModel"
	EntityCreationError = "entityCreationError"
	Unauthorized        = "unauthorized"
)

var errorMessage = map[string]string{
	"invalidUserID":       "invalid user id",
	"internalError":       "an internal error occured",
	"userNotFound":        "user could not be found",
	"invalidBindingModel": "model could not be bound",
	"entityCreationError": "could not create entity",
	"unauthorized":        "an unauthorized access",
}
