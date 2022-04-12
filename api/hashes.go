package api

import (
	"fmt"
	"strconv"
)

// main function to dispatch
func hashes(args []string) (string, error) {
	if len(args) >= 1 {
		switch args[0] {
		case "download":
			args = args[1:]
			// try converting to numeric id
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return "ID has to be numeric\n", nil
			}

			args = args[1:]

			if len(args) >= 1 {
				switch args[0] {
				case "cracked":
					return downloadCracked(id, "cracked")
				case "plain":
					return downloadCracked(id, "plain")
				default:
					return hashesHelp(), nil
				}
			}
		case "help":
			return hashesHelp(), nil
		default:
			return hashesHelp(), nil
		}
	}
	return hashesHelp(), nil
}

func hashesHelp() string {
	return fmt.Sprint(`
Available commands for 'hashes' are:

  download <id> [cracked, plain]
    - cracked downloads a file with a unique list of cracked passwords [hash:password]
    - plain downloads a file with non unique cracked credentials [username:passwords]

`)
}
