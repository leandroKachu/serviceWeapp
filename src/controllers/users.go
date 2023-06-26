package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		fmt.Println(err)
		responses.JSON(w, http.StatusBadRequest, responses.ErrorApi{Err: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5000/createuser", "application/json", bytes.NewBuffer(user))
	if err != nil {
		log.Fatal(err)
		responses.JSON(w, http.StatusInternalServerError, responses.ErrorApi{Err: err.Error()})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.StatusCodeError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
