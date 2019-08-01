# 02-https-server

[HOME](../readme.md)

## certificate generation

If you want to use the certificates with this repository, then you dont need to generate the certificates yourself. If you want to generate a new root, client and server certificate then you can do so by using the utility called minica.
More details on minica can be found [here](https://github.com/jsha/minica).

To generate the root, server and client certificate + private keys follow the below steps
- install the [Go tools](https://golang.org/dl/) and set up your `$GOPATH`.
- install minica using `go get github.com/jsha/minica`.
- go to a specified directory of you choice
- create server certificate by running `minica --domains server-cert`
- if you are running it for the first time, it will generate 4 files.
    - minica.pem (root certificate)
    - minica-key.pem (private key for root)
    - server-cert/cert.pem (certificate for domain "server-cert", signed by root certificate's public key)
    - server-cert/key.pem (private key for domain "server-cert")
- create client certificate by running `minica --domains client-cert`. It will generate 2 new files
    - client-cert/cert.pem (certificate for domain "client-cert", signed by root certificate's public key)
    - client-cert/key.pem (private key for domain "client-cert")

If you generate your own certificates, replace all 6 files in the directory [00-certificates](../00-certificates). The directory naming is very obvious, put your root key and certificate under `00-certificates`,
client-cert files under `00-certificates/client` and serv-cert files under `00-certificates/server`

## setup domains in your local machine

setup your domains as in informed in [Home](../readme.md) and come back here.

## run server

- from project root, navigate to the server directory `cd 02-https-server/server`.
- run the server `go run server.go`

## run client
- in another terminal window, navigate `cd 02-https-server`
- run the client `go run client.go`
- the requests should fail. The whole idea is to demonstrate, why the requests are failing.

**Important**: Make sure to run the server and the client from their respective directories.