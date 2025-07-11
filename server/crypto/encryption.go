package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type AESCrypto struct {
	key []byte
}

func NewAESCrypto(key []byte) *AESCrypto {
	return &AESCrypto{key: key}
}
func (ac *AESCrypto) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(ac.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := aesgcm.Seal(nonce, nonce, plaintext, nil)
	return ciphertext, nil
}

func (ac *AESCrypto) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(ac.key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceSize := aesgcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
  }
  nonce, ct := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesgcm.Open(nil, nonce, ct, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

