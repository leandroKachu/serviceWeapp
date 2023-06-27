package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webapp/responses"
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

	response, err := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(loginUser))
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorApi{Err: err.Error()})
		return
	}

	token, _ := ioutil.ReadAll(response.Body)
	err = json.Unmarshal([]byte(token), &jsonResponse)
	if err != nil {
		fmt.Println("Erro ao decodificar o JSON:", err)
		return
	}

	id := jsonResponse.ID
	tokenValue := jsonResponse.Token
	fmt.Println(id)
	fmt.Println(tokenValue)

}
