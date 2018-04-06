package config

import (
	"encoding/json"
	"log"
	"os"
)

type DbServer struct {
	Schema   string `json:"schema"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	User     string `json:"user"`
	Weight   int    `json:"weight"`
	Read     bool   `json:"read"`
	Write    bool   `json:"write"`
}

type DbInfo struct {
	Type      string     `json:"type"`
	DbServers []DbServer `json:"servers"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type LogConfig struct {
	Access string `json:"access"`
	Error  string `json:"error"`
	Info   string `json:"info"`
}

type LimitConfig struct {
	Request            []int
	RefreshTokenExpire []int
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Config struct {
	Database     DbInfo        `json:"database"`
	Server       ServerConfig  `json:"server"`
	Log          LogConfig     `json:"log"`
	ReidsConfigs []RedisConfig `json:"redisconfigs"`
	Limitation   LimitConfig   `json:"limit"`
	Dev          bool          `json:"dev"`
	Addr         string        `json:"addr"`
	RootUrl      string        `json:"root_url"`
}

var GlobalDefaultConfig Config
var IsDev bool
var HasInit bool

func LoadConfig(filename string) (*Config, error) {
	configer, err := LoadConfigFile(filename)
	if err != nil {
		return nil, err
	}
	HasInit = true
	GlobalDefaultConfig = *configer
	return &GlobalDefaultConfig, nil
}

func LoadConfigFile(filename string) (*Config, error) {
	fd, err := os.Open(filename)
	if err != nil {
		log.Fatalf("open file %s faild , error %v", filename, err)
		return nil, err
	}
	configer := new(Config)
	decloder := json.NewDecoder(fd)
	err = decloder.Decode(configer)
	if err != nil {
		log.Fatalf("decode json file %s faild , error %v", filename, err)
		return nil, err
	}
	return configer, nil
}

func GetConfig() *Config {
	return &GlobalDefaultConfig
}

func GetIsInit() bool {
	return HasInit
}

func GetDatabase() *DbInfo {
	return &GlobalDefaultConfig.Database
}
func GetServer() *ServerConfig {
	return &GlobalDefaultConfig.Server
}
func GetLog() *LogConfig {
	return &GlobalDefaultConfig.Log
}
func GetRedisConfig() *[]RedisConfig {
	return &GlobalDefaultConfig.ReidsConfigs
}
func GetLimitation() *LimitConfig {
	return &GlobalDefaultConfig.Limitation
}

func GetDev() bool {
	return GlobalDefaultConfig.Dev
}

func GetAddr() string {
	return GlobalDefaultConfig.Addr
}

func GetRootUrl() string {
	return GlobalDefaultConfig.RootUrl
}
