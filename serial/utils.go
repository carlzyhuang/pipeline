package serial

import (
	"log"
	"runtime/debug"
)

func RecoverFromPanic(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %+v\nStack Trace:\n%s", err, debug.Stack())
			go RecoverFromPanic(fn)
		}
	}()

	fn()
}

func RecoverGo(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("panic: %+v\nStack Trace:\n%s", err, debug.Stack())
		}
	}()

	fn()
}
