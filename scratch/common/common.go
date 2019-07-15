package common

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"errors"
	"log"
	"math/big"
	"time"
)

func DecryptMessage(b []byte, p *rsa.PrivateKey) []byte {
	d, err := rsa.DecryptPKCS1v15(rand.Reader, p, b)
	if err != nil {
		log.Fatal(err)
	}
	PrettyPrint(string(d))
	return d
}

func EncryptMessage(message string, pub rsa.PublicKey) []byte {
	b := []byte(message)
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, &pub, b)
	if err != nil {
		log.Fatal(err)
	}
	PrettyPrint(cipherText)
	return cipherText
}

func GeneratePKey() *rsa.PrivateKey {
	// 256 for simplification. for better encryption use 2048 or 4096 bytes
	// this size also determines the length of message that can be encrypted
	pKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatal(err)
	}
	PrettyPrint(pKey)
	return pKey
}

func PrettyPrint(i interface{}) {
	data, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", data)
}

func GenerateHash(message string, p *rsa.PrivateKey) []byte {
	h := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, p, crypto.SHA256, h[:])
	if err != nil {
		log.Fatal(err)
	}

	PrettyPrint(signature)
	return signature
}

func VerifySignature(message, signature []byte, pub *rsa.PublicKey) bool {
	h := sha256.Sum256(message)
	if err := rsa.VerifyPKCS1v15(pub, crypto.SHA256, h[:], signature); err != nil {
		return false
	}
	return true
}

func CertTemplate() (*x509.Certificate, error) {
	// generate a random serial number (a real cert authority would have some logic behind this)
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, errors.New("failed to generate serial number: " + err.Error())
	}

	tmpl := x509.Certificate{
		SerialNumber:          serialNumber,
		Subject:               pkix.Name{Organization: []string{"Yhat, Inc."}},
		SignatureAlgorithm:    x509.SHA256WithRSA,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour), // valid for an hour
		BasicConstraintsValid: true,
	}
	return &tmpl, nil
}
