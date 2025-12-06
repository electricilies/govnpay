## ![](/static/shared/icon/code_gm_grey_24dp.svg) Documentation [¶](#section-documentation "Go to Documentation")

### Overview [¶](#pkg-overview "Go to Overview")

Package crypto collects common cryptographic constants.

### Index [¶](#pkg-index "Go to Index")

* [func RegisterHash(h Hash, f func() hash.Hash)](#RegisterHash)
* [func SignMessage(signer Signer, rand io.Reader, msg []byte, opts
  SignerOpts) (signature []byte, err error)](#SignMessage)
* [type Decrypter](#Decrypter)
* [type DecrypterOpts](#DecrypterOpts)
* [type Hash](#Hash)
* + [func (h Hash) Available() bool](#Hash.Available)
  + [func (h Hash) HashFunc() Hash](#Hash.HashFunc)
  + [func (h Hash) New() hash.Hash](#Hash.New)
  + [func (h Hash) Size() int](#Hash.Size)
  + [func (h Hash) String() string](#Hash.String)
* [type MessageSigner](#MessageSigner)
* [type PrivateKey](#PrivateKey)
* [type PublicKey](#PublicKey)
* [type Signer](#Signer)
* [type SignerOpts](#SignerOpts)

### Constants [¶](#pkg-constants "Go to Constants")

This section is empty.

### Variables [¶](#pkg-variables "Go to Variables")

This section is empty.

### Functions [¶](#pkg-functions "Go to Functions")

#### func [RegisterHash](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=145) [¶](#RegisterHash "Go to RegisterHash")

```
func RegisterHash(h Hash, f func() hash.Hash)
```

RegisterHash registers a function that returns a new instance of the
given hash function. This is intended to be called from the init
function in packages that implement hash functions.

#### func [SignMessage](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=245) [¶](#SignMessage "Go to SignMessage") added in go1.25.0

```
func SignMessage(signer Signer, rand io.Reader, msg []byte, opts SignerOpts) (signature []byte, err error)
```

SignMessage signs msg with signer. If signer implements
[MessageSigner](#MessageSigner),
[MessageSigner.SignMessage] is called directly. Otherwise, msg is
hashed with opts.HashFunc() and signed with [Signer.Sign].

### Types [¶](#pkg-types "Go to Types")

#### type [Decrypter](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=229) [¶](#Decrypter "Go to Decrypter") added in go1.5

```
type Decrypter interface {
	// Public returns the public key corresponding to the opaque,
	// private key.
	Public() PublicKey

	// Decrypt decrypts msg. The opts argument should be appropriate for
	// the primitive used. See the documentation in each implementation for
	// details.
	Decrypt(rand io.Reader, msg []byte, opts DecrypterOpts) (plaintext []byte, err error)
}
```

Decrypter is an interface for an opaque private key that can be used
for asymmetric decryption operations. An example would be an RSA key
kept in a hardware module.

#### type [DecrypterOpts](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=240) [¶](#DecrypterOpts "Go to DecrypterOpts") added in go1.5

```
type DecrypterOpts any
```

#### type [Hash](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=16) [¶](#Hash "Go to Hash")

```
type Hash uint
```

Hash identifies a cryptographic hash function that is implemented in
another package.

```
const (
	MD4         Hash = 1 + iota // import golang.org/x/crypto/md4
	MD5                         // import crypto/md5
	SHA1                        // import crypto/sha1
	SHA224                      // import crypto/sha256
	SHA256                      // import crypto/sha256
	SHA384                      // import crypto/sha512
	SHA512                      // import crypto/sha512
	MD5SHA1                     // no implementation; MD5+SHA1 used for TLS RSA
	RIPEMD160                   // import golang.org/x/crypto/ripemd160
	SHA3_224                    // import golang.org/x/crypto/sha3
	SHA3_256                    // import golang.org/x/crypto/sha3
	SHA3_384                    // import golang.org/x/crypto/sha3
	SHA3_512                    // import golang.org/x/crypto/sha3
	SHA512_224                  // import crypto/sha512
	SHA512_256                  // import crypto/sha512
	BLAKE2s_256                 // import golang.org/x/crypto/blake2s
	BLAKE2b_256                 // import golang.org/x/crypto/blake2b
	BLAKE2b_384                 // import golang.org/x/crypto/blake2b
	BLAKE2b_512                 // import golang.org/x/crypto/blake2b

)
```

#### func (Hash) [Available](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=138) [¶](#Hash.Available "Go to Hash.Available")

```
func (h Hash) Available() bool
```

Available reports whether the given hash function is linked into
the binary.

#### func (Hash) [HashFunc](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=19) [¶](#Hash.HashFunc "Go to Hash.HashFunc") added in go1.4

```
func (h Hash) HashFunc() Hash
```

HashFunc simply returns the value of h so that
[Hash](#Hash) implements
[SignerOpts](#SignerOpts).

#### func (Hash) [New](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=127) [¶](#Hash.New "Go to Hash.New")

```
func (h Hash) New() hash.Hash
```

New returns a new hash.Hash calculating the given hash function.
New panics if the hash function is not linked into the binary.

#### func (Hash) [Size](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=116) [¶](#Hash.Size "Go to Hash.Size")

```
func (h Hash) Size() int
```

Size returns the length, in bytes, of a digest resulting from the
given hash function. It doesn't require that the hash function in
question be linked into the program.

#### func (Hash) [String](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=23) [¶](#Hash.String "Go to Hash.String") added in go1.15

```
func (h Hash) String() string
```

#### type [MessageSigner](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=213) [¶](#MessageSigner "Go to MessageSigner") added in go1.25.0

```
type MessageSigner interface {
	Signer
	SignMessage(rand io.Reader, msg []byte, opts SignerOpts) (signature []byte, err error)
}
```

MessageSigner is an interface for an opaque private key that can be
used for signing operations where the message is not pre-hashed by
the caller. It is a superset of the Signer interface so that it can
be passed to APIs which accept Signer, which may try to do an
interface upgrade.

MessageSigner.SignMessage and MessageSigner.Sign should produce the
same result given the same opts. In particular,
MessageSigner.SignMessage should only accept a zero opts.HashFunc if
the Signer would also accept messages which are not pre-hashed.

Implementations which do not provide the pre-hashed Sign API should
implement Signer.Sign by always returning an error.

#### type [PrivateKey](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=176) [¶](#PrivateKey "Go to PrivateKey")

```
type PrivateKey any
```

PrivateKey represents a private key using an unspecified algorithm.

Although this type is an empty interface for backwards compatibility
reasons, all private key types in the standard library implement the
following interface

```
interface{
    Public() crypto.PublicKey
    Equal(x crypto.PrivateKey) bool
}
```

as well as purpose-specific interfaces such as
[Signer](#Signer) and [Decrypter](#Decrypter),
which can be used for increased type safety within applications.

#### type [PublicKey](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=162) [¶](#PublicKey "Go to PublicKey") added in go1.2

```
type PublicKey any
```

PublicKey represents a public key using an unspecified algorithm.

Although this type is an empty interface for backwards compatibility
reasons, all public key types in the standard library implement the
following interface

```
interface{
    Equal(x crypto.PublicKey) bool
}
```

which can be used for increased type safety within applications.

#### type [Signer](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=180) [¶](#Signer "Go to Signer") added in go1.4

```
type Signer interface {
	// Public returns the public key corresponding to the opaque,
	// private key.
	Public() PublicKey

	// Sign signs digest with the private key, possibly using entropy from
	// rand. For an RSA key, the resulting signature should be either a
	// PKCS #1 v1.5 or PSS signature (as indicated by opts). For an (EC)DSA
	// key, it should be a DER-serialised, ASN.1 signature structure.
	//
	// Hash implements the SignerOpts interface and, in most cases, one can
	// simply pass in the hash function used as opts. Sign may also attempt
	// to type assert opts to other types in order to obtain algorithm
	// specific values. See the documentation in each package for details.
	//
	// Note that when a signature of a hash of a larger message is needed,
	// the caller is responsible for hashing the larger message and passing
	// the hash (as digest) and the hash function (as opts) to Sign.
	Sign(rand io.Reader, digest []byte, opts SignerOpts) (signature []byte, err error)
}
```

Signer is an interface for an opaque private key that can be used
for signing operations. For example, an RSA key kept in a hardware
module.

#### type [SignerOpts](https://cs.opensource.google/go/go/%2B/go1.25.5%3Asrc/crypto/crypto.go;l=219) [¶](#SignerOpts "Go to SignerOpts") added in go1.4

```
type SignerOpts interface {
	// HashFunc returns an identifier for the hash function used to produce
	// the message passed to Signer.Sign, or else zero to indicate that no
	// hashing was done.
	HashFunc() Hash
}
```

SignerOpts contains options for signing with a
[Signer](#Signer).