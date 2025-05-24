package menu

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/yaml.v3"
)

const Tab = "\t"

type Option struct {
	Code       string   `yaml:"code"`
	Label      string   `yaml:"label"`
	Command    string   `yaml:"command"`
	SubOptions []Option `yaml:"suboptions"`
}

type Config struct {
	Prompt  string   `yaml:"prompt"`
	Options []Option `yaml:"options"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) CodeForLabel(label string) (string, bool) {
	for _, opt := range c.Options {
		if opt.Label == label {
			return opt.Code, true
		}
	}
	return "", false
}

func (c *Config) CommandForCode(code string) (string, bool) {
	for _, opt := range c.Options {
		if opt.Code == code {
			return opt.Command, true
		}
	}
	return "", false
}

func ShowMenu(cfg *Config) (Option, error) {
	options := cfg.Options
	prompt := cfg.Prompt

	for {
		var labels []string
		codeMap := make(map[string]Option)

		for _, opt := range options {
			label := fmt.Sprintf("%s\t%s", opt.Code, opt.Label)
			labels = append(labels, label)
			codeMap[opt.Code] = opt
		}

		cmd := exec.Command("rofi", "-dmenu", "-p", prompt)
		cmd.Stdin = strings.NewReader(strings.Join(labels, "\n"))
		out, err := cmd.Output()
		if err != nil {
			return Option{}, err
		}

		choice := strings.TrimSpace(string(out))
		if choice == "" {
			return Option{}, fmt.Errorf("no selection made")
		}

		code := strings.Fields(choice)[0]
		selected, ok := codeMap[code]
		if !ok {
			return Option{}, fmt.Errorf("invalid selection")
		}

		if len(selected.SubOptions) > 0 {
			// Recurse deeper into suboptions
			options = selected.SubOptions
			prompt = selected.Label
			continue
		}

		// Return the final selected option
		return selected, nil
	}
}
