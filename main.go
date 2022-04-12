package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/patrickhener/go-cj/api"
	"github.com/patrickhener/go-cj/config"
)

const gocjVersion = "v0.0.1"

func init() {
	version := flag.Bool("v", false, "version information")

	flag.Parse()

	if *version {
		fmt.Printf("go-cj version is: %s\n", gocjVersion)
		os.Exit(1)
	}
}

func main() {
	var err error

	config.GoCJConfig = config.Config{}
	fmt.Println("[*] Checking for config file")
	err = config.GoCJConfig.SaveConfig()
	if err != nil {
		fmt.Println("[-] Config file related error")
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("[*] Loading config file")
	config.GoCJConfig, err = config.LoadConfig()
	if err != nil {
		fmt.Println("[-] Config loading related error")
		fmt.Println(err)
		os.Exit(1)
	}

	// Api loop
	fmt.Println("[*] Starting command loop - type 'help' for information on how to use it")
	api := api.New()
	err = api.Loop()
	if err != nil {
		panic(err)
	}
}
