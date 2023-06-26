package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorApi struct {
	Err string
}

func JSON(w http.ResponseWriter, statusCode int, datas interface{}) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(datas); err != nil {
		log.Fatal(err)
	}
}

func StatusCodeError(w http.ResponseWriter, r *http.Response) {

	var erro ErrorApi

	json.NewDecoder(r.Body).Decode(&erro)

	JSON(w, r.StatusCode, erro)

}
