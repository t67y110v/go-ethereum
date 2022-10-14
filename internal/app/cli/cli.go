package cli

import (
	pt "github.com/c-bata/go-prompt"
)

type Cli struct {
	cli *pt.Prompt
}

func newCli() *Cli {
	c := &Cli{
		cli: pt.New(Executor,
			completer,
			pt.OptionTitle("Interactive shell client"),
			pt.OptionPrefix("->>"),
			pt.OptionInputTextColor(pt.Purple),
			pt.OptionInputBGColor(pt.DarkGray)),
	}

	return c
}
