package api

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
)

// main function to dispatch
func dispatchHashes(input string, l *readline.Instance) error {
	args := strings.Split(input, " ")
	if len(args) >= 1 {
		switch args[0] {
		case "download":
			args = args[1:]
			// try converting to numeric id
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("%s", "ID must be numeric")
			}

			args = args[1:]

			if len(args) >= 1 {
				switch args[0] {
				case "cracked":
					return downloadCracked(id, "cracked")
				case "plain":
					return downloadCracked(id, "plain")
				}
			}
		default:
		}
	}

	hashUsage(l.Stderr())
	return nil
}

var hashCompleter = readline.NewPrefixCompleter(
	readline.PcItem("download",
		readline.PcItemDynamic(getAllSessionIDs(),
			readline.PcItem("cracked"), readline.PcItem("plain"))),
)

func hashUsage(w io.Writer) {
	_, _ = io.WriteString(w, "Hashes commands:\n")
	_, _ = io.WriteString(w, hashCompleter.Tree("    "))
	_, _ = io.WriteString(w, "\n")
}
