package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL = ""
	Port   = 0
)

// var APIURL = ""

// var Port = 0
var HashKey = []byte("")
var BlockKey = []byte("")

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Port, erro = strconv.Atoi(os.Getenv("PORT"))

	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("APIURL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
	fmt.Println(HashKey)

}
