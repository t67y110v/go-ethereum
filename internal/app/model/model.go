package model

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
)

var iv = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const secretKey string = "abc&1*~#^2^#s0^=)^^7%b34"

func encodeBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func decodeBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

type EthWallet struct {
	PrivateKey        string
	PublicKey         string
	PublicAddres      string
	Password          string
	EncryptedPassword string
}

func (e *EthWallet) BeforeCreate() error {
	if len(e.Password) > 0 {
		bv := []byte(e.Password)
		hasher := sha256.New()
		hasher.Write(bv)
		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		enc, err := Encrypt(sha, secretKey)
		if err != nil {
			return err
		}
		e.EncryptedPassword = enc
	}
	return nil
}

func (e *EthWallet) ComparePassword(password string) bool {
	bv := []byte(e.Password)
	hasher := sha256.New()
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	p, err := Encrypt(sha, secretKey)
	if err != nil {
		return false
	}
	if p == e.EncryptedPassword {
		return true
	}
	return false
}

func Encrypt(text, secretKey string) (string, error) {
	bv := []byte(text)
	hasher := sha256.New()
	hasher.Write(bv)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(sha)
	cfb := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encodeBase64(cipherText), nil
}

func Decrypt(text, secretKey string) (string, error) {

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	cipherText := decodeBase64(text)
	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
