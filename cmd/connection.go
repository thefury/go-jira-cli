package cmd

import (
	"github.com/spf13/viper"
)

type ConnectionSettings struct {
	BaseUrl  string
	UserName string
	Token    string
}

func LoadConnection() ConnectionSettings {
	settings := viper.AllSettings()

	connection := ConnectionSettings{
		BaseUrl:  settings["url"].(string),
		UserName: settings["username"].(string),
		Token:    settings["token"].(string),
	}

	return connection
}
