package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config 配置信息
type Config struct {
	Gin     Gin      `yaml:"gin"`
	Token   string   `yaml:"token"`
	Servers []Server `yaml:"servers"`
}

// Gin Gin 配置信息
type Gin struct {
	Mode string `yaml:"mode"`
	Port string `yaml:"port"`
}

// Server 服务器密钥信息
type Server struct {
	Name       string `yaml:"name"`
	InstanceID string `yaml:"instanceID"`
	Region     string `yaml:"region"`
	SecretID   string `yaml:"secretID"`
	SecretKey  string `yaml:"secretKey"`
	RuleTag    string `yaml:"ruleTag"`
}

// Get 获取配置信息
func Get(path string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse config file: %w", err)
	}
	return cfg, nil
}
