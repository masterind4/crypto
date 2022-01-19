


Certif Exo1:

```bash
# Rootca
openssl ecparam -name prime256v1 -genkey -noout -out ca.key
openssl ec -in ca.key -pubout -out ca.pub
openssl req -x509 -new -nodes -key ca.key -sha256 -days 1024 -out rootCA.pem
# Client
openssl ecparam -name prime256v1 -genkey -noout -out cert.key
openssl ec -in cert.key -pubout -out cert.pub
# CSR
openssl req -new -key cert.key -out cert.csr
# Signature
openssl x509 -req -in cert.csr -CA rootCA.pem -CAkey ca.key -CAcreateserial -days 256 -sha256 -out cert.pem -extfile ca.conf -extensions req_ext

# curl
curl -v https://127.0.0.1:4443 -k
curl -v https://127.0.0.1:4443 --cacert rootCA.pem
```
