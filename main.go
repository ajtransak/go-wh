package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	webhookData := "eyJhbGciOiJIUzI1NiJ9.eyJ3ZWJob29rRGF0YSI6eyJpZCI6Ijli.o5gzCsdYOEw8w-M-nUYX5GOwEL8sZyBk0MCU9SwJYgY"

	secretKey := "eyJhbGciOiJIUzI1NiI.eyJBUElfS0VZIjoiZDc5NjcxYjE3MzIxMjkwNDF9.HZVtK3VsiIb2z-J5jLPih_YA"

	token, err := jwt.Parse(webhookData, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		log.Fatalf("Error parsing JWT: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		file, err := os.Create("decoded_jwt_data.json")
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		jsonData, err := json.MarshalIndent(claims, "", "  ")
		if err != nil {
			log.Fatalf("Failed to convert claims to JSON: %v", err)
		}

		if _, err := file.Write(jsonData); err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}

		fmt.Println("Decoded JWT data saved to decoded_jwt_data.json")
	} else {
		log.Println("Token claims are not of expected type jwt.MapClaims")
	}
}
