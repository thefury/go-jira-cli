package cmd

import (
	"encoding/json"
)

type Formatter interface {
	Format() string
}

type JSONFormatter struct {
	Data interface{}
}

func (self *JSONFormatter) Format() string {

	jsonData, err := json.MarshalIndent(self.Data, "", "    ")
	if err != nil {
		panic(err)
		// TODO remove all panics
	}

	return string(jsonData)
}
