package cookies

import (
	"fmt"
	"net/http"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

var s *securecookie.SecureCookie

func ConfigCookie() {
	s = securecookie.New(config.HashKey, config.BlockKey)
}

func Save(w http.ResponseWriter, ID, Token string) error {
	data := map[string]string{
		"id":    ID,
		"token": Token,
	}

	fmt.Println(data)

	dadaEncoded, err := s.Encode("data", data)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Dados",
		Value:    dadaEncoded,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}
