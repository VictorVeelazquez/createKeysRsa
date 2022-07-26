package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	key, err := rsa.GenerateKey(rand.Reader, 128)
	if err != nil {
		fmt.Println(err)
	}

	// Genera la clave privada
	pkcs1PrivateKey := x509.MarshalPKCS1PrivateKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: pkcs1PrivateKey,
	}
	// Escribe en el archivo
	file, err := os.Create("private.pem")
	pem.Encode(file, block)

	// Genera la clave pública
	PublicKey := &key.PublicKey
	// La clave pública se genera a partir de la clave privada
	pkixPublicKey, err := x509.MarshalPKIXPublicKey(PublicKey)
	block1 := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pkixPublicKey,
	}
	file2, err := os.Create("public.pem")
	encode := pem.Encode(file2, block1)
	fmt.Println(encode)

}
