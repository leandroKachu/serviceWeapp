package controllers

import (
	"fmt"
	"net/http"
)

func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome"))
	fmt.Println("cheguei")
}
