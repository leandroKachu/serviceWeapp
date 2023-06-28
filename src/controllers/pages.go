package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/responses"
	"webapp/src/config"
	"webapp/src/model"
	"webapp/src/requests"
	"webapp/src/utils"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "login.html", nil)
}

func LoadViewCreate(w http.ResponseWriter, r *http.Request) {
	utils.RunTemplate(w, "formUser.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)

	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	fmt.Println(response.StatusCode)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{Err: err.Error()})
		return
	}

	if response.StatusCode >= 400 {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{Err: err.Error()})
		return
	}

	var posts []model.Post

	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{Err: err.Error()})
		return
	}

	utils.RunTemplate(w, "home.html", posts)
}
