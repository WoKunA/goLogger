package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type LogEngine struct {
	Config           *LogConfig `yaml:"Config"`
	nowPiece         int
	loopClosedSignal chan bool
	Logger           *log.Logger
	logFile          *os.File
}

func NewLogEngine(config *LogConfig) *LogEngine {
	logEngine := &LogEngine{
		nowPiece: 0,
		Config:   config,
	}
	logEngine.Config.PathEnsure()

	fmt.Println(config.LoopLogFile)
	//是否循环
	if config.LoopLogFile {
		//初始化文件描述符和 Logger
		//init file fd and Logger
		logEngine.SetLoopFileName()
		logEngine.LoadLogFileAndLogger()
		go logEngine.OpenLoopLogFile()
	} else {
		logEngine.LoadLogFileAndLogger()
	}
	return logEngine
}

func (l *LogEngine) CloseLoop() {
	l.loopClosedSignal <- true
}

func (l *LogEngine) OpenLoopLogFile() {
	for {
		time.Sleep(l.Config.Timer)

		select {
		case <-l.loopClosedSignal:
			log.Println("Loop Log Mode Closed !")
			break
		default:
		}

		l.nowPiece = (l.nowPiece + 1) % l.Config.Piece
		l.SetLoopFileName()
		l.LoadLogFileAndLogger()
	}
}

func (l *LogEngine) GetNowPiece() int {
	return l.nowPiece
}

func (l *LogEngine) SetLoopFileName() {
	if len(l.Config.FileName) < 4 {
		err := errors.New("length of log filename wrong or log filename wrong")
		panic(err)
	}
	midfileName := l.Config.FileName
	if midfileName[len(midfileName)-4:len(midfileName)] != ".log" {
		err := errors.New("log file suffix wrong")
		panic(err)
	}
	l.Config.FileName = midfileName[0:len(midfileName)-4] + strconv.Itoa(l.nowPiece) + ".log"
}

func (l *LogEngine) LoadLogFileAndLogger() {
	if l.logFile != nil {
		l.logFile.Close()
	}

	err := l.Config.PathEnsure()
	if err != nil {
		panic(err)
	}

	l.logFile, err = os.OpenFile(l.Config.FileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	l.Logger = log.New(l.logFile, "", 0)
}
