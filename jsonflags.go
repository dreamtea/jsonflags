package jsonflags

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type parser struct {
	flagSet         *flag.FlagSet
	defined         map[string]bool
	configPath      string
	configMustExist bool
}

func (p *parser) getDefinedFlags() {
	p.flagSet.Visit(func(f *flag.Flag) {
		p.defined[f.Name] = true
	})
}

func (p *parser) getConfigPath() bool {
	configFlag := p.flagSet.Lookup("config")
	if configFlag == nil {
		return false
	}
	p.configMustExist = false
	if p.defined["config"] {
		p.configPath = configFlag.Value.String()
		p.configMustExist = true
	} else {
		p.configPath = configFlag.DefValue
	}
	if len(p.configPath) != 0 {
		return true
	}
	p.configMustExist = false
	return false
}

func (p *parser) readJsonConfig() error {
	configFile, err := os.Open(p.configPath)
	if err != nil {
		if err == os.ErrNotExist && !p.configMustExist {
			err = nil
		}
		return err
	}
	defer configFile.Close()

	data := make(map[string]interface{})
	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(&data); err != nil {
		return err
	}
	for key, value := range data {
		key = strings.ToLower(key)
		if !p.defined[key] {
			valueAsString := fmt.Sprintf("%v", value)
			err := p.flagSet.Set(key, valueAsString)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ParseFlagSet(flagSet *flag.FlagSet, args []string) error {
	err := flagSet.Parse(args)
	if err != nil {
		return err
	}
	parser := parser{
		flagSet: flagSet,
		defined: make(map[string]bool, 16),
	}
	parser.getDefinedFlags()
	if parser.getConfigPath() {
		err = parser.readJsonConfig()
	}
	return err
}

func Parse() error {
	return ParseFlagSet(flag.CommandLine, os.Args[1:])
}
