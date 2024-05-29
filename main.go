package main

import (
	"net/http"
	"os"
	"time"

	"github.com/MarioGalvaoWoohoo/api-go-gin/database"
	"github.com/MarioGalvaoWoohoo/api-go-gin/routes"
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", protectedHandler)

	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}

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

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Token ausente", http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	if token.Valid {
		w.Write([]byte("Acesso permitido a rota protegida!"))
	} else {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
	}
}
