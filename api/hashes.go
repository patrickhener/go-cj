package api

import (
	"fmt"
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

	usage(l.Stderr())
	return nil
}
