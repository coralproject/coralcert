package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"strings"
)

// keyIDLength represents the length of characters a keyID is
// when using the GenerateKeyID method.
const keyIDLength = 8

// GenerateKeyID will compute a SHA1-hash of the public key der bytes
// and use the first keyIDLength characters as the key id (kid) to
// return.
func GenerateKeyID(publicKeyDER []byte) (string, error) {

	// Use a SHA1 hash to create a unique identifier for the public key.
	h := sha1.New()
	h.Write(publicKeyDER)

	// Encode the SHA1 sum using a base64 encoding.
	hash := base64.StdEncoding.EncodeToString(h.Sum(nil))

	// Guard against trying to access something not possible (out of bounds).
	if keyIDLength > len(hash) {
		return "", errors.New("key id length not possible, sha1 hashed public key not long enough")
	}

	return hash[:keyIDLength], nil
}

// MarshalPEMBlock will marshal the PEM block in a way that will
// allow marshaling in a JSON payload.
func MarshalPEMBlock(block *pem.Block) string {

	// Encode the block to memory.
	encodedPEM := pem.EncodeToMemory(block)

	// Replace all the newlines with an escaped newline character, this
	// is needed for the JSON parsing.
	return strings.Replace(string(encodedPEM), "\n", "\\n", -1)
}

// MarshalSecret implements the MarshalJSON interface for secrets.
func MarshalSecret(keyID string, pub, pvt *pem.Block) ([]byte, error) {

	// Error if the key id was not found.
	if keyID == "" {
		return nil, errors.New("cannot marshal a secret without a key id")
	}

	// Error if the public key was not found.
	if pub == nil {
		return nil, errors.New("cannot marshal a secret without a public key")
	}

	// Marshal the public key using PEM encoding.
	publicKeyPEM := MarshalPEMBlock(pub)

	// Marshal the private key using PEM encoding if it was provided.
	var privateKeyPEM string
	if pvt != nil {
		privateKeyPEM = MarshalPEMBlock(pvt)
	}

	return json.Marshal(struct {
		KeyID   string `json:"kid"`
		Private string `json:"private,omitempty"`
		Public  string `json:"public"`
	}{
		KeyID:   keyID,
		Private: privateKeyPEM,
		Public:  publicKeyPEM,
	})
}
