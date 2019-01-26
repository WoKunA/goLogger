package logger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Template func(v interface{}) string

type Logger struct {
	Template  Template
	LogEngine *LogEngine `yaml:"LogEngine"`
}

var DefaultLog = NewLog(fmt.Sprintf("%s/logger.yml", os.Getenv("PWD")))

func NewLog(configPath string) *Logger {
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	logger := &Logger{}
	err = yaml.Unmarshal(yamlFile, &logger)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	logger.LogEngine = NewLogEngine(logger.LogEngine.Config)
	logger.Template = LoadToLogger
	return logger
}

func (l *Logger) SetTemplete(fn func(v interface{}) string) {
	l.Template = fn
}

func (l *Logger) Log(v interface{}) {
	l.LogEngine.Logger.Printf("%s", LoadToLogger(v))
}

func LoadToLogger(v interface{}) string {
	TransformJson, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(TransformJson)
}
