package cli

import (
	"io/ioutil"
	"log"

	pt "github.com/c-bata/go-prompt"
)

func completer(d pt.Document) []pt.Suggest {

	var s []pt.Suggest
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
				Text: "signIn", Description: "signin your wallet",
			},
		}
	case "r", "re", "rel", "relo":
		s = []pt.Suggest{
			{
				Text: "reload", Description: "reloading wallet data",
			},
		}
	default:

		s = genWalletList("./wallet")
	}
	return pt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)

}

func genWalletList(path string) []pt.Suggest {
	var s []pt.Suggest
	files, err := ioutil.ReadDir("./wallet")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		s = append(s, pt.Suggest{Text: f.Name(), Description: "Enter your password"})
		//fmt.Println(f.Name())
	}

	//fmt.Println(s)
	return s
}
