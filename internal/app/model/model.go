package model

import (
	"crypto/aes"
	"crypto/cipher"
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
		enc, err := Encrypt(e.Password, secretKey)
		if err != nil {
			return err
		}
		e.EncryptedPassword = enc
	}
	return nil
}

func (e *EthWallet) ComparePassword(password string) bool {
	p, err := Encrypt(password, secretKey)
	if err != nil {
		return false
	}
	if p == e.EncryptedPassword {
		return true
	}
	return false
}

// Encrypt methWalletod is to encrypt or hide any classified text
func Encrypt(text, secretKey string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
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
