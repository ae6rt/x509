package main

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Print useful certificate fields without openssl-x509 fuss.
func main() {
	certFile := flag.String("cert", "-", "File containing one PEM encoded X.509 certificate.  Defaults to standard input.")
	flag.Parse()

	var f io.Reader
	var err error

	switch *certFile {
	case "":
		log.Fatalf("Need a cert file\n")
	case "-":
		f = os.Stdin
	default:
		f, err = os.Open(*certFile)
		if err != nil {
			log.Fatal(err)
		}
	}

	pemBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(pemBytes)
	if block == nil || block.Type != "CERTIFICATE" {
		log.Fatal("failed to decode PEM block containing certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Subject:\t%s\n", cert.Subject.CommonName)
	fmt.Printf("Issuer:\t\t%s\n", cert.Issuer.CommonName)
	fmt.Printf("Subject Key:\t%s\n", strings.ToUpper(hex.EncodeToString(cert.SubjectKeyId)))
	fmt.Printf("Authority Key:\t%s\n", strings.ToUpper(hex.EncodeToString(cert.AuthorityKeyId)))
}
