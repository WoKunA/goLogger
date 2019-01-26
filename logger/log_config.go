package logger

import (
	"fmt"
	"os"
	"time"
)

type LogConfig struct {
	Piece       int           `yaml:"Piece"`
	Timer       time.Duration `yaml:"Timer"`
	FileName    string        `yaml:"FileName"`
	FileFlag    int           `yaml:"FileFlag"`
	LoopLogFile bool          `yaml:"LoopLogFile"`
}

func (l *LogConfig) PathEnsure() (err error) {
	if l.FileName == "" {
		l.FileName = fmt.Sprintf("%s/data/log/Daily.log", os.Getenv("PWD"))
	}
	index := 0
	for i, char := range l.FileName {
		if char == '/' && i != 0 {
			index = i
			path := l.FileName[0:index]
			err = l.CreateDir(path)
			if err != nil {
				return
			}
		}
	}
	return
}

func (l *LogConfig) CreateDir(path string) (err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}
