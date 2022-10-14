package eth

import (
	"ethgo/internal/app/model"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func NewWallet(password string) *model.EthWallet {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	ew := &model.EthWallet{
		Password: password,
	}
	if err := ew.BeforeCreate(); err != nil {
		log.Fatal(err)
	}

	a, err := key.NewAccount(ew.EncryptedPassword)

	if err != nil {
		log.Fatal(err)
	}

	ew.PublicAddres = a.Address.String()
	return ew
}
