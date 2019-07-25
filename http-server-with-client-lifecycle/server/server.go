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
	CertPath            string = "../../certificates/server/cert.pem"
	KeyPath             string = "../../certificates/server/key.pem"
	RootCertificatePath string = "../../certificates/minica.pem"
)

func main() {
	mux := http.NewServeMux()

	// add an endpoint
	mux.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "i am protected")
	})
	log.Println("starting server")

	rootCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	rootCAPool := x509.NewCertPool()
	rootCAPool.AppendCertsFromPEM(rootCA)
	log.Println("RootCA loaded")

	s := &http.Server{
		Handler: mux,
		Addr:    ":8080",
		TLSConfig: &tls.Config{
			ClientCAs:  rootCAPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
			GetCertificate: func(info *tls.ClientHelloInfo) (certificate *tls.Certificate, e error) {
				log.Println("client requested certificate")

				c, err := tls.LoadX509KeyPair(CertPath, KeyPath)
				if err != nil {
					fmt.Printf("Error loading key pair: %v\n", err)
					return nil, err
				}
				return &c, nil
			},
			VerifyPeerCertificate: func(rawCerts [][]byte, chains [][]*x509.Certificate) error {
				if len(chains) > 0 {
					fmt.Println("Verified certificate chain from peer:")

					for _, v := range chains {
						// fmt.Printf("Chain %d:\n", j)
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
