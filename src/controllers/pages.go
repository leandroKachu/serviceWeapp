package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "login.html", nil)
}

func LoadViewCreate(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "formUser.html", nil)
}
