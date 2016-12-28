package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {

	rootPrivateKey, rootCertificate := createCA()
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	csrTemplate := x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:         "domain.com",
			Country:            []string{"DE"},
			Organization:       []string{"Company Ltd"},
			OrganizationalUnit: []string{"IT"},
		},
	}
	csrBytes, _ := x509.CreateCertificateRequest(rand.Reader, &csrTemplate, privateKey)
	pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csrBytes})

	csr, _ := x509.ParseCertificateRequest(csrBytes)
	template := csrToTemplate(csr)

	certificateBytes, err := x509.CreateCertificate(rand.Reader,
		template,
		rootCertificate,
		&privateKey.PublicKey,
		rootPrivateKey)

	if err != nil {
		panic(err)
	}

	pem.Encode(os.Stdout, &pem.Block{Type: "CERTIFICATE", Bytes: certificateBytes})

	privateKeyPemBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	bytes := pem.EncodeToMemory(privateKeyPemBlock)

	fmt.Println(string(bytes))

	encryptedPrivateKeyBlock, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", bytes, []byte("hugo"), x509.PEMCipherAES256)
	if err != nil {
		panic(err)
	}

	pem.Encode(os.Stdout, encryptedPrivateKeyBlock)
	regainedKey, _ := x509.DecryptPEMBlock(encryptedPrivateKeyBlock, []byte("hugo"))
	fmt.Println(string(regainedKey))
}

func createCA() (*rsa.PrivateKey, *x509.Certificate) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	serial, _ := rand.Int(rand.Reader, (&big.Int{}).Exp(big.NewInt(2), big.NewInt(159), nil))
	now := time.Now()

	template := x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			CommonName:         "snippet_ca",
			Country:            []string{"DE"},
			Organization:       []string{"Snippet Ltd"},
			OrganizationalUnit: []string{"IT"},
		},
		NotBefore:             now.Add(-10 * time.Minute).UTC(),
		NotAfter:              now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000).UTC(),
		BasicConstraintsValid: true,
		IsCA:        true,
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
	}
	certificateBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}
	certificate, err := x509.ParseCertificate(certificateBytes)
	if err != nil {
		panic(err)
	}

	return privateKey, certificate
}

func csrToTemplate(csr *x509.CertificateRequest) *x509.Certificate {
	now := time.Now()

	serial, _ := rand.Int(rand.Reader, (&big.Int{}).Exp(big.NewInt(2), big.NewInt(159), nil))

	return &x509.Certificate{
		SerialNumber:          serial,
		Subject:               csr.Subject,
		NotBefore:             now.Add(-10 * time.Minute).UTC(),
		NotAfter:              now.Add(60 * 60 * 24 * 365 * 1000 * 1000 * 1000).UTC(),
		BasicConstraintsValid: true,
		IsCA: false,
	}
}
