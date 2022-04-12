package api

import (
	"fmt"
	"strconv"
)

// main function to dispatch
func sessions(args []string) (string, error) {
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
						return "ID has to be numeric\n", nil
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
					return "ID has to be numeric\n", nil
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
		case "execute":
			args = args[1:]
			if len(args) >= 1 {
				// Implement
			}
		case "help":
			return sessionHelp(), nil
		default:
			return sessionHelp(), nil

		}
	}
	return sessionHelp(), nil
}

func sessionHelp() string {
	return fmt.Sprint(`
Available commands for 'sessions' are:

  get all
  get <id>
  set <id> termination <YYYY-MM-DD-HH:MM>
  set <id> notification [true, false]
  execute <id> [start,stop,pause,rebuild,restore]

`)
}
