package cli

import (
	"fmt"
)

func Start() {
	fmt.Println("DeNet Test Task Go-cli-ethereum")
	fmt.Println("Write `exit` to exit the program")
	fmt.Println("If you have an account, write `reload` and select it from the list, and enter the password to access the wallet")
	fmt.Println("If you do not have an account, enter the newWallet and your password")
	defer fmt.Println("Thx for using this CLI, Bye!")
	cli := newCli()

	cli.cli.Run()
	return
}
