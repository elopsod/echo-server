# Root CA
openssl genrsa -out rootCA.key 4096
openssl req -new -x509 -sha256 -key rootCA.key -subj "/C=CA/ST=APP/L=LOCAL/O=ORG/OU=UNIT/CN=ROOT" -out rootCA.crt -days 36500

# Certificate
openssl genrsa -out server.key 4096
openssl req -new -key server.key -subj "/C=CA/ST=APP/L=LOCAL/O=ORG/OU=UNIT/CN=default" -out server.csr
openssl x509 -req  -in server.csr -CA rootCA.crt -CAkey rootCA.key -CAcreateserial -out server.crt -days 36500 -sha256 -extfile \
    <(printf "subjectAltName=DNS:grpcs.local,DNS:localhost\nkeyUsage=critical,digitalSignature,keyEncipherment\nextendedKeyUsage=serverAuth,clientAuth\nbasicConstraints=critical,CA:FALSE\nauthorityKeyIdentifier=keyid,issuer")
