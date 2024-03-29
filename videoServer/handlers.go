package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"videoServer/api/dbops"
	"videoServer/api/defs"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create user handler")
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, defs.ErrorDBError)
		return
	}

	//id := session.GenerateNewSessionId(ubody.Username)
	//su := &defs.SignedUp{Success: true, SessionId: id}

	//if resp, err := json.Marshal(su); err != nil {
	//	sendErrorResponse(w, defs.ErrorInternalFaults)
	//	return
	//} else {
	//	sendNormalResponse(w, string(resp), 201)
	//}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create user handler")
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
