package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"webapp/responses"
	"webapp/src/config"
	"webapp/src/cookies"
)

func LogUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var jsonResponse struct {
		ID    int64  `json:"id"`
		Token string `json:"token"`
	}

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

	token, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal([]byte(token), &jsonResponse)

	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{Err: err.Error()})
		return
	}

	idStr := strconv.FormatInt(jsonResponse.ID, 10)
	responses.JSON(w, response.StatusCode, nil)

	if erro := cookies.Save(w, idStr, jsonResponse.Token); erro != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, responses.ErrorApi{Err: err.Error()})
		return
	}

}
