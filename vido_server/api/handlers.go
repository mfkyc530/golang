package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"vido_server/api/dbops"
	"vido_server/api/defs"
	"vido_server/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res, ubody); err != nil{
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.UserName, ubody.Pwd); err != nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.UserName)
	su := &defs.SignedUp{Success:true, SessionId:id}
	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}else {
		sendNormalResponse(w, string(resp), 201)
	}
}


func Login (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	uname := p.ByName("user_name")
	io.WriteString(w, uname)

}