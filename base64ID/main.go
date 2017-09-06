package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

const prefixUser string = "USR_"

func main() {

	a := "bar"
	switch a {
	case "foo":
		fmt.Println("This is ok")
		break
	default:
		log.Fatal("Not Implemented")
	}

}

func generateBase64ID(size int, prefix string) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", err

	}

	encodedID := base64.URLEncoding.EncodeToString(b)
	var fullID string
	var idTaken bool

	switch prefix {
	case prefixUser:
		fullID = prefixUser + encodedID
		idTaken = userIdTaken(prefixedId)
		break
	default:
		log.Fatal("Prefix Not Implemented")

	}

	if idTaken {
		return generateBase64ID(size, prefix)

	}
	return fullID, nil
}

// Generate a random base64 url encoded string
func generateBase64ID(size int) (string, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	encoded := base64.URLEncoding.EncodeToString(b)
	return encoded, nil
}
