package controllers

import (
	"encoding/json"
	"github.com/aliwert/go-jwt-example/models"
	u "github.com/aliwert/go-jwt-example/utils"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}
	resp := account.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := models.Login(account.Email, account.Password)
	u.Respond(w, resp)
}
