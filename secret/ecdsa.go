package secret

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
)

// NewECDSA will generate a ECDSA key using the crypto/rand.Reader source
// with a length of bitSize.
func NewECDSA(curve elliptic.Curve) (*ECDSA, error) {
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	return &ECDSA{
		Private: private,
		Public:  private.Public(),
	}, nil
}

// ECDSA provides capabilities around marshaling a
// ECDSA key pair for the Coral Project's Talk.
type ECDSA struct {
	Private *ecdsa.PrivateKey
	Public  crypto.PublicKey
}

// MarshalJSON implements the MarshalJSON interface.
func (s ECDSA) MarshalJSON() ([]byte, error) {

	// Generate the public key from the private ECDSA key.
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
		Type:  "ECDSA PUBLIC KEY",
		Bytes: publicKeyDER,
	}

	// Generate the kid parameter from the public key.
	keyID, err := GenerateKeyID(publicKeyDER)
	if err != nil {
		return nil, err
	}

	// Prepare the ASN.1 DER encoded form in a PEM block if it exists.
	var privateKeyPEM *pem.Block
	if s.Private != nil {

		// Encode the public key to a DER-encoded ASN.1 format.
		privateKeyDER, err := x509.MarshalECPrivateKey(s.Private)
		if err != nil {
			return nil, err
		}

		privateKeyPEM = &pem.Block{
			Type:  "ECDSA PRIVATE KEY",
			Bytes: privateKeyDER,
		}
	}

	// Marshal using the Secret marshaler.
	return Marshal(keyID, publicKeyPEM, privateKeyPEM)
}
