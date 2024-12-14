package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"go-rest-api/config"
	"go-rest-api/types"
	"go-rest-api/utils"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"strconv"
	"time"
)

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(secret string, userID int) (string, error) {
	expiration := time.Second * time.Duration(config.EnvVars.JWTExpirationInSeconds)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   strconv.Itoa(userID),
		"expireAt": time.Now().Add(expiration).Unix(),
	})

	tokenStr, err := t.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func WithJWTAuth(handler http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := getTokenFromRequest(r)

		token, err := validateToken(tokenString)
		if err != nil {
			log.Printf("failed to validate token %v", err)
			permissionDenied(w)
			return
		}

		if !token.Valid {
			log.Printf("failed to validate token %v", err)
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		str := claims["userID"].(string)

		idInt, _ := strconv.Atoi(str)

		u, err := store.GetUserByID(idInt)
		if err != nil {
			log.Printf("failed to validate token %v", err)
			permissionDenied(w)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, u.ID)
		r = r.WithContext(ctx)
		handler(w, r)
	}
}

func GetUserIDFromContext(ctx context.Context) int {
	userID, ok := ctx.Value(UserKey).(int)
	if !ok {
		return -1
	}
	return userID
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteError(w, http.StatusForbidden, fmt.Errorf("permission denied"))
}

func getTokenFromRequest(r *http.Request) string {
	authToken := r.Header.Get("Authorization")
	return authToken
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.EnvVars.JWTSecret), nil
	})
}
