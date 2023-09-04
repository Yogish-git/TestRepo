package signuplogin

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/yogish-git/JeeraApp_Project/pkg/config"
	"github.com/yogish-git/JeeraApp_Project/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

var db = config.GetDB()

var jwtSecret = []byte("your-secret-key") // Replace with a secure secret key

func generateJWTToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AuthorizeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized - token string empty", http.StatusUnauthorized)
			return
		}

		// Split the "Bearer <token>" header to get the actual token
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		actualToken := tokenParts[1]

		token, err := jwt.Parse(actualToken, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized- parsing error", http.StatusUnauthorized)
			return
		}

		// Extract user claims from the token
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Now you can access claims like claims["username"]

		next.ServeHTTP(w, r)
	})
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := DecodeJSON(r.Body, &newUser)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	newUser.Password = string(hashedPassword)
	db.Create(&newUser)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "User registered successfully")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := DecodeJSON(r.Body, &loginData)
	if err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	var user models.User
	db.Where("username = ?", loginData.Username).First(&user)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//jwt code begins
	isAuthenticated := true

	if isAuthenticated {
		token, err := generateJWTToken(loginData.Username)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"token": token}
		EncodeJSON(w, response)
		return
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
	//jwt code ends

	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintln(w, "Login successful🎉")
}

func DecodeJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func EncodeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
