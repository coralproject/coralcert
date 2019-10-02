package main

import (
	"crypto/elliptic"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/coralproject/coralcert/secret"
	"github.com/pkg/errors"
)

func run(bitSize int, curve, secretType string) error {
	switch secretType {
	case "rsa":
		if bitSize < 1024 {
			return errors.Errorf("-bit_size=%d too small, must be at least 1024 bits", bitSize)
		}

		// Generate a secret.
		s, err := secret.NewRSA(bitSize)
		if err != nil {
			return errors.Wrap(err, "couldn't generate the rsa secret")
		}

		// Marshal the key as JSON.
		if err := json.NewEncoder(os.Stdout).Encode(s); err != nil {
			return errors.Wrap(err, "couldn't marshal the rsa secret")
		}

		return nil
	case "ecdsa":
		var eCurve elliptic.Curve
		switch curve {
		case "P256":
			eCurve = elliptic.P256()
		case "P384":
			eCurve = elliptic.P384()
		case "P521":
			eCurve = elliptic.P521()
		default:
			return errors.Errorf("-curve=%s not supported, supports P256 P384 or P521 curves", curve)
		}

		s, err := secret.NewECDSA(eCurve)
		if err != nil {
			return errors.Wrap(err, "couldn't generate the ecdsa secret")
		}

		// Marshal the key as JSON.
		if err := json.NewEncoder(os.Stdout).Encode(s); err != nil {
			return errors.Wrap(err, "couldn't marshal the ecdsa secret")
		}

		return nil
	default:
		return errors.Errorf("-type=%s not supported, use either rsa or ecdsa", secretType)
	}
}

func main() {
	bitSize := flag.Int("bit_size", 2048, "bit size of generated keys if using -type=rsa, minimum 1024")
	curve := flag.String("curve", "P256", "elliptic curve to use if using -type=ecdsa, supports P256 P384 or P521 curves")
	secretType := flag.String("type", "ecdsa", "type of secret to generate, either ecdsa or rsa")

	flag.Parse()

	if err := run(*bitSize, *curve, *secretType); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
