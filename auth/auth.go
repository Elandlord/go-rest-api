package auth

import (
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"mentechmedia.nl/rest-api/app/handler"
)

func Authenticate(writer http.ResponseWriter, request *http.Request) {

	name := request.FormValue("application_name")
	password := request.FormValue("password")

	if len(name) == 0 || len(password) == 0 {
		handler.RespondError(writer, http.StatusBadRequest, "Please provide name and password to obtain the token")
		return
	}

	if name == "eric" && password == "test" {

		token, err := getToken(name)

		if err != nil {
			handler.RespondError(writer, http.StatusNotFound, err.Error())
			return
		}

		writer.Header().Set("Authorization", "Bearer "+token)
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("Token: " + token))

	} else {
		handler.RespondError(writer, http.StatusUnauthorized, "Name and password do not match")
		return
	}
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, response *http.Request) {

		tokenString := response.Header.Get("Authorization")

		if len(tokenString) == 0 {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Missing Authorization Header"))
			return
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := verifyToken(tokenString)

		if err != nil {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}

		name := claims.(jwt.MapClaims)["name"].(string)
		role := claims.(jwt.MapClaims)["role"].(string)

		response.Header.Set("name", name)
		response.Header.Set("role", role)

		next.ServeHTTP(writer, response)
	})
}
