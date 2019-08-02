# 04-https-server-mtls

[HOME](../readme.md)

- navigate to server directory `cd 04-https-server-mtls/server`
- run the server `go run server.go`
- on another terminal, run the client `go run client/client.go`

In this case we have protected both the client and the server by validating the certificates received from the opposite party. Check out the code documentation for more details