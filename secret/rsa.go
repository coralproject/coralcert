package secret

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// NewRSA will generate a RSA key using the crypto/rand.Reader source
// with a length of bitSize.
func NewRSA(bitSize int) (*RSA, error) {
	private, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	return &RSA{
		Private: private,
		Public:  private.Public(),
	}, nil
}

// RSA provides capabilities around marshaling a
// RSA key pair for the Coral Project's Talk.
type RSA struct {
	Private *rsa.PrivateKey
	Public  crypto.PublicKey
}

// MarshalJSON implements the MarshalJSON interface.
func (s RSA) MarshalJSON() ([]byte, error) {
	// Generate the public key from the private RSA key.
	var publicKey crypto.PublicKey
	if s.Public != nil {
		publicKey = s.Public
	} else {
		publicKey = s.Private.Public()
	}

	// Encode the public key to a DER-encoded PKIX format.
	publicKeyDER, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	// Prepare the DER-encoded PKIX encoded form in a PEM block.
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyDER,
	}

	// Generate the kid parameter from the public key.
	keyID, err := GenerateKeyID(publicKeyDER)
	if err != nil {
		return nil, err
	}

	// Prepare the ASN.1 DER encoded form in a PEM block.
	var privateKeyPEM *pem.Block
	if s.Private != nil {
		privateKeyPEM = &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(s.Private),
		}
	}

	// Marshal using the Secret marshaler.
	return Marshal(keyID, publicKeyPEM, privateKeyPEM)
}
