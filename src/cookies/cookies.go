package cookies

import (
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

	dadaEncoded, err := s.Encode("MeuCookie", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "MeuCookie",
		Value:    dadaEncoded,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("MeuCookie")
	if err != nil {
		return nil, err
	}

	values := make(map[string]string)

	if err = s.Decode("MeuCookie", cookie.Value, &values); err != nil {
		return nil, err
	}

	return values, nil

}
