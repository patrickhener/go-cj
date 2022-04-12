package api

import (
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

var Mode string = ""

func modeUsage(w io.Writer) {
	_, _ = io.WriteString(w, "Basic commands:\n")
	_, _ = io.WriteString(w, modeCompleter.Tree("    "))
	_, _ = io.WriteString(w, "\n")
}

var modeCompleter = readline.NewPrefixCompleter(
	readline.PcItem("sessions"),
	readline.PcItem("hashes"),
	readline.PcItem("back"),
	readline.PcItem("help"),
	readline.PcItem("exit"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block Ctrl+Z feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func Run() {
	startConfig := &readline.Config{
		Prompt:          "\033[31m\033[1;35mgo-cj »\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    modeCompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	}
	l, err := readline.NewEx(startConfig)

	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.SetOutput(l.Stderr())

	for {
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		line = strings.TrimSpace(line)

		// Switch mode
		switch {
		case line == "sessions":
			Mode = "sessions"
			l.SetPrompt("\033[31m\033[1;35mgo-cj (sessions) »\033[0m ")
			l.Config.AutoComplete = sessionCompleter
		case line == "hashes":
			Mode = "hashes"
			l.SetPrompt("\033[31m\033[1;35mgo-cj (hashes) »\033[0m ")
			l.Config.AutoComplete = hashCompleter
		case line == "back":
			Mode = ""
			l.SetPrompt(startConfig.Prompt)
			l.Config.AutoComplete = modeCompleter
		case line == "help":
			modeUsage(l.Stderr())
		case line == "exit":
			goto exit
		default:
		}

		// If there is a mode dispatch
		if Mode != "" {
			switch Mode {
			case "sessions":
				if err := dispatchSessions(strings.TrimSpace(line), l); err != nil {
					log.Printf("Error: %+v", err)
				}
			case "hashes":
				if err := dispatchHashes(strings.TrimSpace(line), l); err != nil {
					log.Printf("Error: %+v", err)
				}
			}
		}
	}
exit:
}
