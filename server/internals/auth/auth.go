package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	jwt.RegisteredClaims
	AccessCode string
	UserId     string
}

// CreateJWT this function help to create a new auth token to help authenticate user
func Create(id string, accessCode string) (string, error) {

	// setting the timeout claims for the token to create
	claims := Token{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "findr.com",
			Subject:   "Access Authentication",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		AccessCode: accessCode,
		UserId:     id,
	}

	// create a new uer token for authentication
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		// lg.Error.Println(err.Error())
		return "", err
	}

	// lg.Info.Println("JWT token created")

	return token, nil
}

func Parse(authToken string) (*Token, error) {
	token, err := jwt.ParseWithClaims(authToken, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		// lg.Error.Println("cannot parse auth token : ", err.Error())
		return nil, err
	}

	if claims, ok := token.Claims.(*Token); ok && token.Valid {
		// lg.Info.Println("token parsed is valid ")
		return claims, nil
	}

	return nil, err
}
