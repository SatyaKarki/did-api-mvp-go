package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/mr-tron/base58"
)

type KeyPair struct {
	PublicKey  string
	PrivateKey string
	RawPub     ed25519.PublicKey
	RawPriv    ed25519.PrivateKey
}

// GenerateEd25519KeyPair generates a new Ed25519 key pair and encodes public key in Base58
func GenerateEd25519KeyPair() (*KeyPair, error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	pub58 := base58.Encode(pub)
	//priv64 := base64.StdEncoding.EncodeToString(priv)
	priv64 := base58.Encode(priv)

	return &KeyPair{
		PublicKey:  pub58,
		PrivateKey: priv64,
		RawPub:     pub,
		RawPriv:    priv,
	}, nil
}

func SignMessage(message []byte, priv ed25519.PrivateKey) (string, error) {
	signature := ed25519.Sign(priv, message)
	return base64.StdEncoding.EncodeToString(signature), nil
}

func SignMessageNew(message []byte, base58PrivKey string) (string, error) {
	privKeyBytes, err := base58.Decode(base58PrivKey)
	if err != nil {
		return "", err
	}
	if len(privKeyBytes) != ed25519.PrivateKeySize {
		return "", errors.New("invalid private key length")
	}

	sig := ed25519.Sign(privKeyBytes, message)
	return base58.Encode(sig), nil
}

func VerifyMessage(message []byte, base58PublicKey string, base64Signature string) (bool, error) {
	pubKeyBytes, err := base58.Decode(base58PublicKey)
	if err != nil {
		return false, err
	}
	signature, err := base64.StdEncoding.DecodeString(base64Signature)
	if err != nil {
		return false, err
	}
	return ed25519.Verify(pubKeyBytes, message, signature), nil
}

// VerifySignature checks whether a signature is valid for the given message and public key
func VerifySignature(message []byte, signature []byte, publicKey []byte) (bool, error) {
	if len(publicKey) != ed25519.PublicKeySize {
		return false, errors.New("invalid public key length")
	}

	if len(signature) != ed25519.SignatureSize {
		return false, errors.New("invalid signature length")
	}

	valid := ed25519.Verify(publicKey, message, signature)
	return valid, nil
}
