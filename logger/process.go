package logger

import "log"

func ProcessError(text string, err error) {
	if err != nil {
		log.Println(text, err)
	}
}
