package utils

import (
	"ecom/internal/models"
	"encoding/base64"
	"net/http"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	jwt.Claims
	Role []models.Role `json:"role"`
}

func ExtractBearerTokenFromHeader(r *http.Request) (string, bool) {
	if r == nil || r.Header == nil {
		return "", false
	}
	auth := r.Header.Get("Authorization")
	if len(auth) > 7 && auth[:7] == "Bearer " {
		return auth[7:], true
	}
	return "", false
}

func ExtractBasicAuthCredentials(r *http.Request) (string, string, bool) {
	auth := r.Header.Get("Authorization")
	if len(auth) > 6 && auth[:6] == "Basic " {
		decoded, err := base64.StdEncoding.DecodeString(auth[6:])
		if err != nil {
			return "", "", false
		}
		parts := string(decoded)
		if len(parts) > 0 {
			credentials := string(decoded)
			for i, c := range credentials {
				if c == ':' {
					return credentials[:i], credentials[i+1:], true
				}
			}
			return credentials, "", true
		}
	}
	return "", "", false
}

func GenerateJWTToken(claims *JwtCustomClaims, secretKey string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ParseJWTToken(tokenString string, secretKey string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func ValidateJWTToken(tokenString string, secretKey string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return false, err
	}

	return token.Valid, nil
}
