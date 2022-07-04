package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"runtime"

	"golang.org/x/crypto/ripemd160"

	"github.com/MVRetailManager/MVInventoryChain/logging"
)

const (
	checksumLen = 4
	version     = byte(0x00)
)

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	priKey, err := ecdsa.GenerateKey(curve, rand.Reader)

	if err != nil {
		logging.ErrorLogger.Printf("%v", err)
		runtime.Goexit()
	}

	pubKey := append(priKey.PublicKey.X.Bytes(), priKey.PublicKey.Y.Bytes()...)

	return *priKey, pubKey
}

func NewWallet() *Wallet {
	privateKey, publicKey := NewKeyPair()

	return &Wallet{privateKey, publicKey}
}

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		logging.ErrorLogger.Printf("%v", err)
		runtime.Goexit()
	}

	publicRipMD160Hash := hasher.Sum(nil)

	return publicRipMD160Hash
}

func Checksum(data []byte) []byte {
	fstHash := sha256.Sum256(data)
	sndHash := sha256.Sum256(fstHash[:])

	return sndHash[:checksumLen]
}

func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)

	versionHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionHash)

	fullHash := append(versionHash, checksum...)

	address := Base58Encode(fullHash)

	return address
}

func ValidateAddress(address string) bool {
	fullHash, err := Base58Decode([]byte(address))
	if err != nil {
		return false
	}

	actualChecksum := fullHash[len(fullHash)-checksumLen:]
	version := fullHash[0]
	pubKeyHash := fullHash[1 : len(fullHash)-checksumLen]
	targetChecksum := Checksum(append([]byte{version}, pubKeyHash...))

	return bytes.Compare(actualChecksum, targetChecksum) == 0
}
