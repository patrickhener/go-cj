package api

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
)

// main function to dispatch
func dispatchSessions(input string, l *readline.Instance) error {
	args := strings.Split(input, " ")
	if len(args) >= 1 {
		switch args[0] {
		case "get":
			args = args[1:]
			if len(args) >= 1 {
				switch args[0] {
				case "all":
					// return all sessions
					return getAllSessions()
				default:
					// try converting to numeric id
					id, err := strconv.Atoi(args[0])
					if err != nil {
						return fmt.Errorf("%s", "ID must be numeric")
					}

					return getSpecificSession(id)
				}
			}
		case "set":
			args = args[1:]
			if len(args) >= 1 {
				// try converting to numeric id
				id, err := strconv.Atoi(args[0])
				if err != nil {
					return fmt.Errorf("%s", "ID must be numeric")
				}
				// DEBUG
				fmt.Println(id)

				args = args[1:]
				if len(args) >= 2 {
					switch args[0] {
					case "termination":
						// Implement
					case "notification":
						// Implement
					}
				}
			}
		default:
		}
	}

	usage(l.Stderr())
	return nil
}
