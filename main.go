package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

// create a random key

//	func init() {
//		hashkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
//		blockkey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
//		fmt.Println(blockkey)
//		fmt.Println(hashkey)
//	}
func main() {
	config.Load()
	cookies.ConfigCookie()
	utils.LoadTemplate()
	r := router.Gerar()

	fmt.Println("Starting in port", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
