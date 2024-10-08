package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text: '%s', salt: %s", text, salt)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}

func main() {
	var text = "this is secret"
	fmt.Printf("original : %s\n\n", text)

	var sha = sha1.New()
	sha.Write([]byte(text))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)

	fmt.Println(encryptedString)
	// f4ebfd7a42d9a43a536e2bed9ee4974abf8f8dc8

	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)
	// 929fd8b1e58afca1ebbe30beac3b84e63882ee1a

	var hashed2, salt2 = doHashUsingSalt(text)
	fmt.Printf("hashed 2 : %s\n\n", hashed2)
	// cda603d95286f0aece4b3e1749abe7128a4eed78

	var hashed3, salt3 = doHashUsingSalt(text)
	fmt.Printf("hashed 3 : %s\n\n", hashed3)
	// 9e2b514bca911cb76f7630da50a99d4f4bb200b4

	_, _, _ = salt1, salt2, salt3
}
