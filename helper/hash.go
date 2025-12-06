package helper

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type HashAlgo string

const (
	Sha256     HashAlgo = "SHA256"
	HmacSha512 HashAlgo = "HMAC-SHA512"
	MD5        HashAlgo = "MD5"
)

func ComputeSecureHash(data string, hashAlgo HashAlgo, hashSecret string) string {
	var h hash.Hash

	switch hashAlgo {
	case Sha256:
		h = sha256.New()
	case HmacSha512:
		h = hmac.New(sha512.New, []byte(hashSecret))
	case MD5:
		h = md5.New()
	default:
		return ""
	}

	if hashAlgo == HmacSha512 {
		h.Write([]byte(data))
	}

	if hashAlgo == Sha256 || hashAlgo == MD5 {
		h.Write([]byte(hashSecret + data))
	}

	return hex.EncodeToString(h.Sum(nil))
}

func VerifySecureHash(data string, hashAlgo HashAlgo, hashSecret, expectedHash string) bool {
	computedHash := ComputeSecureHash(data, hashAlgo, hashSecret)
	return hmac.Equal([]byte(computedHash), []byte(expectedHash))
}
