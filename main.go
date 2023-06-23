package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Rodando web")

	utils.LoadTemplate()
	r := router.Gerar()

	fmt.Println(r, "http")
	log.Fatal(http.ListenAndServe(":3000", r))
}
