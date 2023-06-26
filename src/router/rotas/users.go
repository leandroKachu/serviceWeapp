package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

var routerUsers = []Rota{
	{
		URI:        "/create-account",
		Metodo:     http.MethodGet,
		Funcao:     controllers.LoadViewCreate,
		RequerAuth: false,
	},
	{
		URI:        "/send-api",
		Metodo:     http.MethodPost,
		Funcao:     controllers.CreateUser,
		RequerAuth: false,
	},
}
