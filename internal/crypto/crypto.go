package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/curve25519"
	"golang.org/x/crypto/hkdf"
	"io"
	"strings"
)

type Sealed struct {
	Kind []byte
	Key  []byte
	Data []byte
}

func (s Sealed) String() string {
	return fmt.Sprintf("%s:%s:%s", base64.URLEncoding.EncodeToString(s.Kind), base64.URLEncoding.EncodeToString(s.Key), base64.URLEncoding.EncodeToString(s.Data))
}
func ParseSealed(value string) (Sealed, error) {
	parts := strings.Split(value, ":")
	if len(parts) != 3 {
		return Sealed{}, errors.New("invalid sealed")
	}
	kind, err := base64.URLEncoding.DecodeString(parts[0])
	if err != nil {
		return Sealed{}, err
	}
	key, err := base64.URLEncoding.DecodeString(parts[1])
	if err != nil {
		return Sealed{}, err
	}
	data, err := base64.URLEncoding.DecodeString(parts[2])
	if err != nil {
		return Sealed{}, err
	}
	return Sealed{
		Kind: kind,
		Key:  key,
		Data: data,
	}, nil
}
func (sealed Sealed) Open(privateKey []byte) ([]byte, error) {
	switch string(sealed.Kind) {
	case "v1":
		return OpenV1(privateKey, sealed)
	default:
		return nil, errors.New("unsupported sealed version")
	}
}

type Key []byte

func (k Key) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}
func ParseKey(s string) (Key, error) {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return Key{}, fmt.Errorf("failed to parse base64-encoded key: %v", err)
	}
	return b, nil
}

func NewKey() (Key, error) {
	privateKey := make([]byte, curve25519.ScalarSize)
	if _, err := rand.Read(privateKey); err != nil {
		return nil, err
	}
	return privateKey, nil
}
func (k Key) PublicKey() (Key, error) {
	publicKey, err := curve25519.X25519(k, curve25519.Basepoint)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

func SealV1(pubKey []byte, data []byte) (Sealed, error) {

	ephermeral, err := NewKey()
	if err != nil {
		return Sealed{}, err
	}
	ephermeralPublic, err := ephermeral.PublicKey()
	if err != nil {
		return Sealed{}, err
	}

	sharedSecret, err := curve25519.X25519(ephermeralPublic[:], pubKey[:])
	if err != nil {
		return Sealed{}, err
	}

	salt := make([]byte, 0, len(ephermeralPublic)+len(pubKey[:]))
	salt = append(salt, ephermeralPublic[:]...)
	salt = append(salt, pubKey[:]...)
	h := hkdf.New(sha256.New, sharedSecret, salt, []byte("X25519"))
	encryptionKey := make([]byte, chacha20poly1305.KeySize)
	if _, err := io.ReadFull(h, encryptionKey); err != nil {
		return Sealed{}, err
	}

	aesCipher, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return Sealed{}, err
	}
	aesGCM, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return Sealed{}, err
	}
	// since we generate a new ephemeral sender key, we don't need a nonce
	nonce := make([]byte, aesGCM.NonceSize())
	encryptedData := aesGCM.Seal(nil, nonce, data, nil)

	// this is another options we could encrypt with:
	// aead, err := chacha20poly1305.New(encryptionKey)
	//if err != nil {
	//	return Sealed{}, err
	//}
	// since we generate a new ephemeral sender key, we don't need a nonce
	//nonce := make([]byte, chacha20poly1305.NonceSize)
	//encryptedData := aead.Seal(nil, nonce, data, nil)

	return Sealed{
		Kind: []byte("v1"),
		Key:  ephermeralPublic[:],
		Data: encryptedData,
	}, nil
}

func OpenV1(privateKey []byte, sealed Sealed) ([]byte, error) {
	ephemeralKey := sealed.Key
	if len(ephemeralKey) != curve25519.PointSize {
		return nil, errors.New("invalid ephemeral key")
	}

	publicKey, err := curve25519.X25519(privateKey, curve25519.Basepoint)
	if err != nil {
		return nil, fmt.Errorf("invalid key: %w", err)
	}

	sharedSecret, err := curve25519.X25519(ephemeralKey, publicKey[:])
	if err != nil {
		return nil, fmt.Errorf("invalid key: %w", err)
	}

	salt := make([]byte, 0, len(ephemeralKey)+len(publicKey[:]))
	salt = append(salt, ephemeralKey...)
	salt = append(salt, publicKey[:]...)
	h := hkdf.New(sha256.New, sharedSecret, salt, []byte("X25519"))
	encryptionKey := make([]byte, chacha20poly1305.KeySize)
	if _, err := io.ReadFull(h, encryptionKey); err != nil {
		return nil, err
	}

	aesCipher, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	return aesGCM.Open(nil, nonce, sealed.Data, nil)

	// this is another options we could decrypt with:
	//aead, err := chacha20poly1305.New(encryptionKey)
	//if err != nil {
	//	return nil, err
	//}
	//nonce := make([]byte, chacha20poly1305.NonceSize)
	//return aead.Open(nil, nonce, sealed.Data, nil)
}
