package main

import (
	"crypto/elliptic"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/coralproject/coralcert/secret"
)

func main() {

	bitSize := flag.Int("bit_size", 2048, "bit size of generated keys if using -type=rsa, minimum 1024")
	curve := flag.String("curve", "P256", "elliptic curve to use if using -type=ecdsa, supports P256 P384 or P521 curves")
	secretType := flag.String("type", "ecdsa", "type of secret to generate, either ecdsa or rsa")

	flag.Parse()

	switch *secretType {
	case "rsa":
		if *bitSize < 1024 {
			fmt.Fprintf(os.Stderr, "-bit_size=%d too small, must be at least 1024 bits\n", *bitSize)
			os.Exit(1)
		}

		// Generate a secret.
		s, err := secret.NewRSA(*bitSize)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't generate the rsa secret: %s\n", err.Error())
			os.Exit(1)
		}

		// Marshal the key as JSON.
		if err := json.NewEncoder(os.Stdout).Encode(s); err != nil {
			fmt.Fprintf(os.Stderr, "couldn't marshal the rsa secret: %s\n", err.Error())
			os.Exit(1)
		}
	case "ecdsa":
		var eCurve elliptic.Curve
		switch *curve {
		case "P256":
			eCurve = elliptic.P256()
		case "P384":
			eCurve = elliptic.P384()
		case "P521":
			eCurve = elliptic.P521()
		default:
			fmt.Fprintf(os.Stderr, "-curve=%s not supported, supports P256 P384 or P521 curves\n", *curve)
			os.Exit(1)
		}

		s, err := secret.NewECDSA(eCurve)
		if err != nil {
			fmt.Fprintf(os.Stderr, "couldn't generate the ecdsa secret: %s\n", err.Error())
			os.Exit(1)
		}

		// Marshal the key as JSON.
		if err := json.NewEncoder(os.Stdout).Encode(s); err != nil {
			fmt.Fprintf(os.Stderr, "couldn't marshal the ecdsa secret: %s\n", err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "-type=%s not supported, use either rsa or ecdsa\n", *secretType)
		os.Exit(1)
	}
}
