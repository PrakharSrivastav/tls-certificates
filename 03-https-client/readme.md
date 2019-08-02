# 03-http-server
[HOME](../readme.md)

This section explains how to configure client correctly to trust the server certificate using a pool of trusted clients.

Since server's certificate is signed by public key of the root CA, we can cryptographically validate the server certificate.

- To do that, we first create a certificate pool that can hold one or more CA certificates.
- We then read all the CA certificates and load it on the CA certificate pool.

In our case, we only have one root CA certificate so we just load one certificate. In general its a common practice to load a chain of root CAs, intermediate CAs on the client. 