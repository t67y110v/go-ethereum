package cli

import (
	//pt "github.com/c-bata/go-prompt"
	"ethgo/internal/app/eth"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	//"strconv"
	"strings"
)

type WalletList struct {
	Wallet string
}

func Executor(s string) {
	s = strings.TrimSpace(s)
	setCommand := strings.Split(s, " ")
	switch setCommand[0] {
	case "exit":
		os.Exit(0)
		return

	case "newWallet":
		ew := eth.NewWallet(setCommand[1])
		fmt.Printf("A new account has been registered.\nPublic address  : %s\n", ew.PublicAddres)
		return
	case "reload":
		files, err := ioutil.ReadDir("./wallet")
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			fmt.Println(f.Name())
		}

	case "signIn":
		ew := eth.ParseWallet("./wallet/"+setCommand[1], setCommand[2])
		fmt.Printf("Successful login.\n Your wallet data:\n Private key:%s\n Public key:%s\n Public addres:%s\n",
			ew.PrivateKey, ew.PublicKey, ew.PublicAddres)
	default:
		return

	}

}
