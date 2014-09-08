package logger

import (
	"bufio"
	"log"
	"os"
)

var (
	logFileName string
	logChannel  chan string
)

func init() {
	logFileName = "fortinet.log"
	logChannel = make(chan string)

	go func(logChannel <-chan string) {
		for msg := range logChannel {
			f, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
			if err != nil {
				// TODO: If this panics it would break the goroutine, would nothing
				// be logged ever again?
				panic(err)
			}

			w := bufio.NewWriter(f)
			log.SetOutput(w)
			log.Println(msg)

			w.Flush()
			f.Close()
		}
	}(logChannel)
}

func Log(message string) {
	logChannel <- message
}
