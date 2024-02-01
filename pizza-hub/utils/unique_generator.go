package utils

import "crypto/rand"

func UniqueGenerator(length int) string {
	characterSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-"
	characterSetLength := len(characterSet)

	randomBytes := make([]byte, length)
	rand.Read(randomBytes)

	var result string
	for _, b := range randomBytes {
		result += string(characterSet[int(b)%characterSetLength])
	}

	return result
}
