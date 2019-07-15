package main

import (
	"github.com/PrakharSrivastav/test-certificates/scratch/common"
)

const plainText string = "This is the plaintext"


// private key includes the public key
// distribute the public key for encryption
// use the private key to read the message
// ie.
// 1. you generate a privateKey
// 2. share the public key for others to encrypt message
// 3. use the private key to decrypt message
// Confirm ownership : Only you can read the messages
func main() {
	// Generate a key
	keyPRV := common.GeneratePKey()
	keyPUB := keyPRV.PublicKey

	// encrypt the plaintext
	cipherText := common.EncryptMessage(plainText, keyPUB)

	// decrypt the cipherText
	_ = common.DecryptMessage(cipherText, keyPRV)
}

