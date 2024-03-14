package hpsdk

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type SdkConfig struct {
	AccessKey      string `yaml:"access_key"`
	Secret         string `yaml:"secret"`
	EndPoint       string `yaml:"endpoint"`
	RequestTimeout int    `yaml:"request_timeout"`
}

func LoadSdkConfig() (*SdkConfig, error) {
	configPath := "configs/config.yaml"

	configBody, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config content of %v: %v", configPath, err)
	}

	conf := &SdkConfig{}

	decoder := yaml.NewDecoder(bytes.NewReader(configBody))
	decoder.KnownFields(true)
	if err := decoder.Decode(conf); err != nil {
		return nil, fmt.Errorf("could not parse config: %v", err)
	}

	return conf, nil
}
