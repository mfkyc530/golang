package main

import (
	"net/http"
	"vido_server/api/defs"
	"vido_server/api/session"
)

var HEDAER_FIELD_SESSION = "X-Session-Id"
var HEDAER_FIELD_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEDAER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEDAER_FIELD_UNAME, uname)
	return true
}

func ValidateUser (w http.ResponseWriter, r *http.Request) bool{
	uname := r.Header.Get(HEDAER_FIELD_UNAME)
	if len(uname) ==0 {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return false
	}

	return true
}
