package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/responses"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/model"
)

func LogUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	loginUser, err := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ErrorApi{Err: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(loginUser))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorApi{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
	}

	var dataAuth model.DataAuth
	if err := json.NewDecoder(response.Body).Decode(&dataAuth); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{Err: err.Error()})
		return
	}

	if err = cookies.Save(w, dataAuth.ID, dataAuth.Token); err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{Err: err.Error()})
		return
	}

}
