package cli

import (
	"io/ioutil"

	pt "github.com/c-bata/go-prompt"
)

func completer(d pt.Document) []pt.Suggest {

	var s []pt.Suggest
	var err error
	switch d.Text {
	case "n", "ne", "new":
		s = []pt.Suggest{
			{Text: "newWallet", Description: "Register a new wallet"},
		}
	case "e", "ex", "exi", "exit":
		s = []pt.Suggest{
			{
				Text: "exit", Description: "Exit cli",
			},
		}
	case "s", "si", "sign":
		s = []pt.Suggest{
			{
				Text: "signIn", Description: "use your wallet and enter password",
			},
		}
	case "r", "re", "rel", "relo":
		s = []pt.Suggest{
			{
				Text: "reload", Description: "reloading wallet data",
			},
		}
	default:

		s, err = genWalletList("./wallet")
		if err != nil {
			s = []pt.Suggest{
				{
					Text: "newWallet", Description: "There are no registered wallets in the system, create a wallet using the command newWallet",
				},
			}
		}
	}
	return pt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)

}

func genWalletList(path string) ([]pt.Suggest, error) {
	var s []pt.Suggest
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		s = append(s, pt.Suggest{Text: f.Name(), Description: "run signIn and enter your password"})
		//fmt.Println(f.Name())
	}

	//fmt.Println(s)
	return s, nil
}
