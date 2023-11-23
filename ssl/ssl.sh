#!/bin/bash
# Generate a private key (you will be prompted to set a passphrase)
openssl genpkey -algorithm RSA -out server.key

# Generate a self-signed certificate using the private key
openssl req -new -x509 -key server.key -out server.crt -days 365