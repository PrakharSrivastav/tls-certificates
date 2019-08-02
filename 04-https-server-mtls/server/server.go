package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	CertPath            string = "../../00-certificates/server/cert.pem"
	KeyPath             string = "../../00-certificates/server/key.pem"
	RootCertificatePath string = "../../00-certificates/minica.pem"
)

func main() {

	// add an endpoint
	mux := http.NewServeMux()
	mux.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "i am protected")
	})

	// create a certificate pool and load all the CA certificates that you
	// want to validate a client against
	clientCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	clientCAPool := x509.NewCertPool()
	clientCAPool.AppendCertsFromPEM(clientCA)
	log.Println("ClientCA loaded")

	// configure http server with tls configuration
	s := &http.Server{
		Handler: mux,
		Addr:    ":8080",
		TLSConfig: &tls.Config{
			ClientCAs:  clientCAPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
			// Loads the server's certificate and sends it to the client
			GetCertificate: func(info *tls.ClientHelloInfo) (certificate *tls.Certificate, e error) {
				log.Println("client requested certificate")
				c, err := tls.LoadX509KeyPair(CertPath, KeyPath)
				if err != nil {
					fmt.Printf("Error loading server key pair: %v\n", err)
					return nil, err
				}
				return &c, nil
			},
			// Call back function to print client certificate details
			VerifyPeerCertificate: func(rawCerts [][]byte, chains [][]*x509.Certificate) error {
				if len(chains) > 0 {
					fmt.Println("Verified certificate chain from peer:")
					for _, v := range chains {
						for i, cert := range v {
							fmt.Printf("  Cert %d:\n", i)
							fmt.Printf(CertificateInfo(cert))
						}
					}
				}
				return nil
			},
		},
	}

	log.Println("starting server")

	// use server.ListenAndServeTLS instead of http.ListenAndServeTLS
	log.Fatal(s.ListenAndServeTLS("", ""))
}

func CertificateInfo(cert *x509.Certificate) string {
	if cert.Subject.CommonName == cert.Issuer.CommonName {
		return fmt.Sprintf("    Self-signed certificate %v\n", cert.Issuer.CommonName)
	}
	s := fmt.Sprintf("    Subject %v\n", cert.DNSNames)
	s += fmt.Sprintf("    Usage %v\n", cert.ExtKeyUsage)
	s += fmt.Sprintf("    Issued by %s\n", cert.Issuer.CommonName)
	s += fmt.Sprintf("    Issued by %s\n", cert.Issuer.SerialNumber)
	return s
}
