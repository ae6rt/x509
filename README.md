## x509

_x509_ is a tool to print commonly useful fields in an X.509 certificate.

## Build

```bash
$ go install
```

## Run

```bash
$ x509 -h
Usage of x509:
  -cert string
    	File containing one PEM encoded X.509 certificate. (default "-")
```

```bash
$ x509 -cert cert.pem
Subject:       Acme Intermediate CA
Issuer:        Acme Parent 
Subject Key:   61E46CC0893D198755D426F3CCCF334624686208
Authority Key: B023BF3B62F5BAB425CE86CDA6EF249D7108A236
```

```bash
$ cat cert.pem | x509 -cert -
Subject:       Acme Intermediate CA
Issuer:        Acme Parent 
Subject Key:   61E46CC0893D198755D426F3CCCF334624686208
Authority Key: B023BF3B62F5BAB425CE86CDA6EF249D7108A236
```
