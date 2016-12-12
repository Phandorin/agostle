package main

import (
	"io"
	"strings"

	"golang.org/x/crypto/openpgp"
)

// readKeyring reads the keyring, and panics on error
func readKeyring(r io.Reader) openpgp.KeyRing {
	var keyring openpgp.EntityList
	for {
		el, err := openpgp.ReadArmoredKeyRing(r)
		if err != nil {
			if len(keyring) == 0 {
				panic(err)
			}
			break
		}
		keyring = append(keyring, el...)
	}
    return keyring
}

var keyring = readKeyring(strings.NewReader(
`
-----BEGIN PGP PUBLIC KEY BLOCK-----
Name: Producer (overseer-bindiff) <T.Gulacsi+producer@unosoft.hu>

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
knO4ag==
=I2Ne
-----END PGP PUBLIC KEY BLOCK-----
-----BEGIN PGP PUBLIC KEY BLOCK-----
Name: Consumer (overseer-bindiff) <T.Gulacsi+consumer@unosoft.hu>

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
H4PoLg==
=HxkD
-----END PGP PUBLIC KEY BLOCK-----
-----BEGIN PGP PRIVATE KEY BLOCK-----
Name: Consumer (overseer-bindiff) <T.Gulacsi+consumer@unosoft.hu>

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
=o0Ex
-----END PGP PRIVATE KEY BLOCK-----`))
