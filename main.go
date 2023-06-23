package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Rodando web")
	r := router.Gerar()

	fmt.Println(r, "http")
	log.Fatal(http.ListenAndServe(":3000", r))
}
