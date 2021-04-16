package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	ReleaseMode string     `yaml:"release_mode"`
	ProjectName string	   `yaml:"project_name"`
	Version     string     `yaml:"version"`
	ServerPort  int        `yaml:"server_port"`
	LogConfig   *LogConfig `yaml:"log_config"`
	MysqlConf   *MysqlConf `yaml:"mysql_conf"`
}

type LogConfig struct {
	LogLevel  string `yaml:"log_level"`
	LogFormat string `yaml:"log_format"`
	LogPath   string `yaml:"log_path"`
	SaveDay   int    `yaml:"save_day"`
}

type MysqlConf struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

var Conf = new(Config)

func InitConf(cfg string) (err error) {
	var yamlFile []byte
	yamlFile, err = ioutil.ReadFile(cfg)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, Conf)
	if err != nil {
		return err
	}
	return err
}

func GetConf() *Config {
	return Conf
}
