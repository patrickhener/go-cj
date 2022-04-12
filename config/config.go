package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var GoCJConfig Config

type Config struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	TLS     bool   `json:"tls"`
	Auth    string `json:"auth"`
	BaseURI string
}

func (c *Config) SaveConfig() error {
	home := os.Getenv("HOME")
	if _, err := os.Stat(fmt.Sprintf("%s/.config/gocj/config.json", home)); os.IsNotExist(err) {
		// Copy
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("No config file found @ ~/.config/gocj/config.json. Do you want me to copy the example there? Type 'y' to do so.")
		decision, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if strings.ToLower(decision) == "y\n" {
			// read template
			input, err := ioutil.ReadFile("config.example.json")
			if err != nil {
				return fmt.Errorf("error reading in example config: %s", err)
			}

			// make folder path
			err = os.MkdirAll(fmt.Sprintf("%s/.config/gocj", home), os.ModePerm)
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(fmt.Sprintf("%s/.config/gocj/config.json", home), input, 0600)
			if err != nil {
				return fmt.Errorf("error writiting mafia.json @ ~/.config/gocj: %s", err)
			}
			return fmt.Errorf("Config was copied successful to %s. Now configure it and rerun", home+"/.config/gocj/config.json")
		} else {
			fmt.Println("Okay bye.")
			os.Exit(0)
		}

	} else if err != nil {
		// Exit no config
		panic(err)
	}
	return nil
}

func LoadConfig() (Config, error) {
	var cfg Config

	home := os.Getenv("HOME")
	cfile, err := ioutil.ReadFile(fmt.Sprintf("%s/.config/gocj/config.json", home))
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(cfile, &cfg)
	if err != nil {
		return Config{}, err
	}
	var schema string
	if cfg.TLS {
		schema = "https"
	} else {
		schema = "http"
	}
	cfg.BaseURI = fmt.Sprintf("%s://%s:%d/api/v1/", schema, cfg.Host, cfg.Port)

	return cfg, nil
}
