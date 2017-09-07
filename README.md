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
{"kid":"0JtIaawv","private":"-----BEGIN ECDSA PRIVATE KEY-----\\nMIGkAgEBBDBYIaCeQSMwciah85K9KzQDj/9JdJDRdy4hxMfmLnfow9ZugjFCD1Lw\\naZgCcjWAJhygBwYFK4EEACKhZANiAASyuTtHHTJ6dFO+9ke/xVtzXh6LAfjJMQII\\nvb3qCf7wzV/ik6Ev92T+IXOk6Qro08fcDKjPlo6fM7quMvDdUxo5rNJRVAA+0NDz\\nSSOoLwJBpdD76JFn2p7b5HwXH0ZTLRE=\\n-----END ECDSA PRIVATE KEY-----\\n","public":"-----BEGIN ECDSA PUBLIC KEY-----\\nMHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEsrk7Rx0yenRTvvZHv8Vbc14eiwH4yTEC\\nCL296gn+8M1f4pOhL/dk/iFzpOkK6NPH3Ayoz5aOnzO6rjLw3VMaOazSUVQAPtDQ\\n80kjqC8CQaXQ++iRZ9qe2+R8Fx9GUy0R\\n-----END ECDSA PUBLIC KEY-----\\n"}
```

## License

    Copyright 2017 Mozilla Foundation

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

        http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

    See the License for the specific language governing permissions and limitations under the License.
