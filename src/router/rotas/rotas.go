package rotas

import (
	"net/http"
	"webapp/src/middleware"

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

	rotas = append(rotas, routerUsers...)
	rotas = append(rotas, rotaHome...)

	for _, rota := range rotas {

		if rota.RequerAuth {
			router.HandleFunc(rota.URI, middleware.Logger(middleware.Auth(rota.Funcao))).Methods(rota.Metodo)
		} else {
			router.HandleFunc(rota.URI, middleware.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))

	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
