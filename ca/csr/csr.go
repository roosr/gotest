package csr

import (
	"crypto/x509"
	"encoding/base64"
	"strings"

	caErr "github.com/roosr/gotest/ca/error"
)

type Csr struct {
	Country, Organization, OrganizationalUnit []string
	City, State                               []string
	StreetAddress, PostalCode                 []string
	SerialNumber, CommonName                  string
	DNSNames                                  []string
	EmailAddresses                            []string

	SignatureAlgorithm x509.SignatureAlgorithm

	PublicKeyAlgorithm x509.PublicKeyAlgorithm
	PublicKey          interface{}
}

func ParseCsr(csrPem string) (*Csr, error) {

	csrBytes, err := parseCsrPem(csrPem)
	if err != nil {
		return nil, err
	}

	// TODO(roos): does this do checks on the modulus and public expoenent to esnure it positive, etc...
	// TODO(roos): the hash algorithm should be check by the calling app

	var certReq *x509.CertificateRequest
	certReq, err = x509.ParseCertificateRequest(csrBytes)
	if err != nil {
		return nil, caErr.New(caErr.ParseCsr)
	}

	err = certReq.CheckSignature()
	if err != nil {
		return nil, caErr.New(caErr.CsrSignature)
	}

	csr := &Csr{
		CommonName:         certReq.Subject.CommonName,
		SerialNumber:       certReq.Subject.SerialNumber,
		StreetAddress:      certReq.Subject.StreetAddress,
		PostalCode:         certReq.Subject.PostalCode,
		City:               certReq.Subject.Locality,
		State:              certReq.Subject.Province,
		Country:            certReq.Subject.Country,
		Organization:       certReq.Subject.Organization,
		OrganizationalUnit: certReq.Subject.OrganizationalUnit,
		DNSNames:           certReq.DNSNames,
		EmailAddresses:     certReq.EmailAddresses,
		SignatureAlgorithm: certReq.SignatureAlgorithm,
		PublicKeyAlgorithm: certReq.PublicKeyAlgorithm,
		PublicKey:          certReq.PublicKey,
	}

	return csr, nil
}

func parseCsrPem(csrPem string) ([]byte, error) {

	var csrBytes []byte

	// Find the start of the CSR (look past any PEM header).
	startPos := strings.Index(csrPem, "MI")
	if startPos < 0 {
		return csrBytes, caErr.New(caErr.FindCsrBase64)
	}
	csrPem = csrPem[startPos:]

	// Check if we need to trim the PEM trailer.
	startPos = strings.Index(csrPem, "-")
	if startPos > 0 {
		csrPem = csrPem[:startPos]
	}

	var err error
	csrBytes, err = base64.StdEncoding.DecodeString(csrPem)
	if err != nil {
		return csrBytes, caErr.New(caErr.DecodeCsrBase64)
	}

	return csrBytes, nil
}
