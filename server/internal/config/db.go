package config

type Database struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	AccountName string `yaml:"account_name"`
	Password    string `yaml:"password"`
	Database    string `yaml:"database"`
}
