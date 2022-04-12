package api

import (
	"fmt"
	"os"
	"strings"

	"github.com/kisielk/cmd"
)

// Functions
func help(args []string) (string, error) {
	text := fmt.Sprint(`
Available API commands are:

  sessions
  hashes
  hashcat
  mask
  wordlist
  rules

Use '<command> help' to get more information on a specific command.

Program commands:

  help - print help
  exit - terminate api client

`)
	return text, nil
}

func exit(args []string) (string, error) {
	fmt.Println("Thanks for using go-cj, bye!")
	os.Exit(1)
	return "", nil
}

func def(line string) (string, error) {
	line = strings.TrimSuffix(line, "\n")
	return fmt.Sprintf("The command '%s' is either not defined or not implemented yet.\n", line), nil
}

// Loop Instance
func New() *cmd.Cmd {
	var commands map[string]cmd.CmdFn = make(map[string]cmd.CmdFn)
	commands["help"] = help
	commands["sessions"] = sessions
	commands["exit"] = exit

	cli := cmd.New(commands, os.Stdin, os.Stdout)
	cli.Default = def
	cli.Prompt = "go-cj > "

	return cli
}
