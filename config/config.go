package config

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var (
	conf *config
	// Kafka Kafka info
	Kafka kafka
	// Port http port
	Port int
)

// Mysql mysql属性
type kafka struct {
	Host  []string
	Topic string
}

// Config 配置内容
type config struct {
	Kafka    kafka `yaml:"kafka"`
	Port     int   `yaml:"port"`
	original string
}

func init() {
	var err error
	conf, err = LoadFile("altermangetokafka.yml")
	if err != nil {
		log.Fatalf("Fail to parse 'altermangetokafka.yml': %v", err)
	}
	Kafka = conf.Kafka
	Port = conf.Port
}

// Load 加载配置
func Load(s string) (*config, error) {
	cfg := &config{}
	err := yaml.UnmarshalStrict([]byte(s), cfg)
	if err != nil {
		return nil, err
	}
	cfg.original = s
	return cfg, nil
}

// LoadFile parses the given YAML file into a Config.
func LoadFile(filename string) (*config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg, err := Load(string(content))
	if err != nil {
		return nil, fmt.Errorf("parsing YAML file %s: %v", filename, err)
	}
	return cfg, nil
}
