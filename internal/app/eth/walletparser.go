package eth

import (
	"ethgo/internal/app/model"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

const secretKey string = "abc&1*~#^2^#s0^=)^^7%b34"

func ParseWallet(path, password string) *model.EthWallet {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	ep, err := model.Encrypt(password, secretKey)
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, ep)
	//key, err := keystore.DecryptKey(b, password)
	privateData := crypto.FromECDSA(key.PrivateKey)
	privateKey := hexutil.Encode(privateData)
	publicData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	publicKey := hexutil.Encode(publicData)
	publicAddres := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	ew := &model.EthWallet{
		PrivateKey:   privateKey,
		PublicKey:    publicKey,
		PublicAddres: publicAddres,
	}
	return ew
}
