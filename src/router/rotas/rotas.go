package rotas

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI        string
	Metodo     string
	Funcao     func(http.ResponseWriter, *http.Request)
	RequerAuth bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin

	fmt.Println("entrei aqui")
	fmt.Println(rotas)

	for _, rota := range rotas {
		fmt.Println("entrei aqui tambem")

		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return router
}
