package utils

import "log"

func Print(v ...any) {
	log.Default().Print(v...)
}
