package utils

import "log"

func Log(v ...any) {
	log.Default().Print(v...)
}
