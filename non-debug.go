// +build !debug

package debug

// ON reports whether debug mode is enabled.
const ON = false

// SetPrefix set the prefix for debug messages.
func SetPrefix(s string) {
	// do nothing
}

// Print print message to the std logger if debug mode is enabled.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	// do nothing
}

// Printf print message to the std logger if debug mode is enabled.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	// do nothing
}

// Println print message to the std logger if debug mode is enabled.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	// do nothing
}

// Assert print message to the std logger and call panic() when predict is false and debug mode is enabled.
// Do nothing otherwise.
func Assert(predict bool, v ...interface{}) {
	// do nothing
}

// Assertf print message to the std logger and call panic() when predict is false and debug mode is enabled.
// Do nothing otherwise.
func Assertf(predict bool, format string, v ...interface{}) {
	// do nothing
}

// Assertln print message to the std logger and call panic() when predict is false and debug mode is enabled.
// Do nothing otherwise.
func Assertln(predict bool, v ...interface{}) {
	// do nothing
}

func Trace(int) {}
