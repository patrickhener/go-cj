package api

import (
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
)

func usage(w io.Writer) {
	_, _ = io.WriteString(w, "Available commands:\n")
	_, _ = io.WriteString(w, completer.Tree("    "))
}

var completer = readline.NewPrefixCompleter(
	readline.PcItem("sessions",
		readline.PcItem("get",
			readline.PcItem("all"), readline.PcItemDynamic(getAllSessionIDs())),
		readline.PcItem("set",
			readline.PcItemDynamic(getAllSessionIDs(),
				readline.PcItem("termination"), readline.PcItem("notification")))),
	readline.PcItem("hashes",
		readline.PcItem("download",
			readline.PcItem("<id>",
				readline.PcItem("cracked"), readline.PcItem("plain")))),
	readline.PcItem("help"),
	readline.PcItem("exit"),
)

func filterInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}

func Run() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31mgo-cj »\033[0m ",
		HistoryFile:     "/tmp/readline.tmp",
		AutoComplete:    completer,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",

		HistorySearchFold:   true,
		FuncFilterInputRune: filterInput,
	})

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
		switch {
		case strings.HasPrefix(line, "sessions "):
			if err := dispatchSessions(line[9:], l); err != nil {
				log.Printf("Error: %+v", err)
			}
		case strings.HasPrefix(line, "hashes "):
			if err := dispatchHashes(line[7:], l); err != nil {
				log.Printf("Error: %+v", err)
			}

		case line == "help":
			usage(l.Stderr())
		case line == "exit":
			goto exit
		}
	}
exit:
}
