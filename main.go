package main

import (
	"fmt"
	"golang.org/x/crypto/chacha20"
	"os"
)

func main() {
	fmt.Println("Enter a message, key, and nonce:")
	var plaintext string
	fmt.Scanln(&plaintext)

	UserKey := make([]byte, chacha20.KeySize)
	UserNonce := make([]byte, chacha20.NonceSize)

	//TRYING TO FIGURE OUT WHY Scanln DOES NOT WORK
	//if _, err := fmt.Scanln(&UserKey); err != nil {
	//	fmt.Println("Error reading key:", err)
	//	return
	//}
	//if _, err := fmt.Scanln(&UserNonce); err != nil {
	//	fmt.Println("Error reading nonce:", err)
	//	return
	//}
	// (WRONG KEY SIZE, TRYING TO FIGURE OUT WHY IT DOESNT WORK; USING os.Stdin.Read INSTEAD)

	if _, err := os.Stdin.Read(UserKey); err != nil {
		fmt.Println("Error reading key:", err)
		return
	}
	if _, err := os.Stdin.Read(UserNonce); err != nil {
		fmt.Println("Error reading nonce:", err)
		return
	}

	cipher, err := chacha20.NewUnauthenticatedCipher(UserKey, UserNonce)
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		return
	}

	ciphertext := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertext, []byte(plaintext))

	fmt.Printf("Encrypted message: %x\n", ciphertext)

	decipher, err := chacha20.NewUnauthenticatedCipher(UserKey, UserNonce)
	if err != nil {
		fmt.Println("Error creating decipher:", err)
		return
	}

	decryptedPlaintext := make([]byte, len(ciphertext))
	decipher.XORKeyStream(decryptedPlaintext, ciphertext)

	fmt.Println("Decrypted message:", string(decryptedPlaintext))
}
