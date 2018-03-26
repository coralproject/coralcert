# coralcert

This package provides a command line utiltiy, `coralcert`, to assist with certificate
generation for use with the Coral Project's [Talk](https://github.com/coralproject/talk) product.

## Installation

You can install from source by running:

```bash
go get github.com/coralproject/coralcert
```

Or by visiting the [Releases](https://github.com/coralproject/coralcert/releases/latest) page for binary builds.

### Installation Via Homebrew

```
brew install coralproject/stable/coralcert
```

## Usage

This utility when ran will output the JSON encoded certificates in a form that is compatible
with the Coral Project's [Talk](https://github.com/coralproject/talk) product to stdout.

```
Usage of coralcert:
  -bit_size int
    	bit size of generated keys if using -type=rsa, minimum 1024 (default 2048)
  -curve string
    	elliptic curve to use if using -type=ecdsa, supports P256 P384 or P521 curves (default "P256")
  -type string
    	type of secret to generate, either ecdsa or rsa (default "ecdsa")
```

## Example

```bash
$ coralcert -type=ecdsa -curve=P384
{"kid":"jhCOJ+3U","private":"-----BEGIN EC PRIVATE KEY-----\\nMIGkAgEBBDAEiASgHtSHyWMQUH2lIFgYwwJiuOdlG2cRbPyueiV6R1+vaZu24Jo+\\nUPsF3MdrbDWgBwYFK4EEACKhZANiAATH+EextZhmUr8m3P1yEdn7+Y76vOrvIL5n\\nZs8pcua60bMMrp3oEu/Tvk8C/+ULBgrRFGPnKtDLcXRQ767GUoafITW44K7D22pd\\nmijHUrUw7jjKlsYnfGCvSjCyri9GmSQ=\\n-----END EC PRIVATE KEY-----\\n","public":"-----BEGIN PUBLIC KEY-----\\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEx/hHsbWYZlK/Jtz9chHZ+/mO+rzq7yC+\\nZ2bPKXLmutGzDK6d6BLv075PAv/lCwYK0RRj5yrQy3F0UO+uxlKGnyE1uOCuw9tq\\nXZoox1K1MO44ypbGJ3xgr0owsq4vRpkk\\n-----END PUBLIC KEY-----\\n"}
```

## License

    Copyright 2017 Mozilla Foundation

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

    See the License for the specific language governing permissions and limitations under the License.
