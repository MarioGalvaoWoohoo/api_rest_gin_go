package authentication

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("seu-secret-key-aqui")

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Verifique as credenciais (normalmente contra um banco de dados)
	username := "usuario"
	password := "senha"

	// Simples exemplo de autenticação
	if r.FormValue("username") == username && r.FormValue("password") == password {
		// Crie um token JWT
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
			return
		}

		w.Write([]byte(tokenString))
	} else {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
	}
}
