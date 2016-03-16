package logger

import "log"

// ProcessError using as common logger
func ProcessError(text string, err error) {
	if err != nil {
		log.Println(text, err)
	}
}
