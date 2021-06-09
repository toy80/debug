// Package debug implement Assert, Trace and Print[f|ln] functions for debugging purpose.

// +build debug

package debug

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

// ON reports whether debug mode is enabled.
const ON = true

var prefix = `(DEBUG) `

// SetPrefix set the prefix for debug messages.
func SetPrefix(s string) {
	prefix = s
}

// Print print message to the std logger if debug mode is enabled.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	log.Output(2, prefix+fmt.Sprint(v...))
}

// Printf print message to the std logger if debug mode is enabled.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	log.Output(2, prefix+fmt.Sprintf(format, v...))
}

// Println print message to the std logger if debug mode is enabled.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	log.Output(2, prefix+fmt.Sprintln(v...))
}

// Assert print message to the std logger and terminate the program immediately when predict is false and debug mode is enabled.
// Do nothing otherwise.
func Assert(predict bool, v ...interface{}) {
	if predict {
		return
	}
	var s string
	if len(v) == 0 {
		s = "Assert failure."
	} else {
		s = "Assert failure: " + fmt.Sprint(v...)
	}
	log.Output(2, s)
	runtime.Breakpoint()
}

// Assertf print message to the std logger and terminate the program immediately when predict is false and debug mode is enabled.
// Do nothing otherwise.
func Assertf(predict bool, format string, v ...interface{}) {
	if predict {
		return
	}
	var s string
	if len(v) == 0 {
		s = "Assert failure."
	} else {
		s = "Assert failure: " + fmt.Sprintf(format, v...)
	}
	log.Output(2, s)
	runtime.Breakpoint()
}

// Assertln print message to the std logger and terminate the program immediately when predict is false and debug mode is enabled.
// Do nothing otherwise.
func Assertln(predict bool, v ...interface{}) {
	if predict {
		return
	}
	var s string
	if len(v) == 0 {
		s = "Assert failure."
	} else {
		s = "Assert failure: " + fmt.Sprintln(v...)
	}
	log.Output(2, s)
	runtime.Breakpoint()
}

func Trace(deep int) {
	if deep < 0 {
		deep = 0
	}
	for i := 0; i <= deep; i++ {
		pc, _, _, ok := runtime.Caller(i + 1)
		if !ok {
			break
		}
		var fn string
		if pf := runtime.FuncForPC(pc); pf != nil {
			fn = pf.Name()
			if pos := strings.LastIndexByte(fn, '/'); pos != -1 {
				fn = fn[pos+1:]
			}
		}
		if i == 0 {
			Printf("call %s", fn)
		} else {
			Printf("  by %s", fn)
		}
	}
}

// func printStack() {
// 	// print stack an terminate the program immediately
// 	buf := debug.Stack()
// 	for i := 0; i < 8; i++ {
// 		if pos := bytes.IndexByte(buf, '\n'); pos != -1 {
// 			buf = buf[pos+1:]
// 		}
// 	}
// 	os.Stderr.Write(buf)
// }
