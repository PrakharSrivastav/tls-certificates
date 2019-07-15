package main

import (
	"log"

	"github.com/PrakharSrivastav/test-certificates/scratch/common"
)

const plainText string = "This is the plaintext"

// in order to confirm that message from the server has not been evesdropped
// and altered in transit. the server can also create a signature and send it to client
// client can then verify if the signature matched the hash of the message
func main() {
	keyPRV := common.GeneratePKey()
	keyPUB := keyPRV.PublicKey
	// this signature is sent to the client
	signature := common.GenerateHash(plainText, keyPRV)


	// using the private key to verify message
	log.Println(common.VerifySignature([]byte("WrongMessage"), signature, &keyPUB))
	log.Println(common.VerifySignature([]byte("AnotherWrongMessage"), signature, &keyPUB))
	log.Println(common.VerifySignature([]byte(plainText), signature, &keyPUB))
}
