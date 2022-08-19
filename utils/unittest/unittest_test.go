package unittest

import (
	"fmt"
	"os"
	"testing"
)

// TestCrashTest_ErrorMessage tests that CrashTest() can check a function that crashed without any messages.
func TestCrashTest_NoMessage(t *testing.T) {
	f := func(t *testing.T) {
		crash_NoMessage()
	}

	CrashTest(t, f, "")
}

// TestCrashTest_ErrorMessage tests that CrashTest() can read standard messages from stdout before a crash.
func TestCrashTest_ErrorMessage(t *testing.T) {
	f := func(t *testing.T) {
		crash_ErrorMessage()
	}
	CrashTest(t, f, "about to crash")
}

// TestCrashTest_ExitCode tests that CrashTest() catches exit codes other than 1.
func TestCrashTest_ExitCode(t *testing.T) {
	f := func(t *testing.T) {
		crash_ExitCode()
	}
	CrashTest(t, f, "exit code != 1")
}

// TestCrashTest_Logger tests that CrashTest() can read fatal logger messages from stdout before a crash. This test
// assumes that the logger uses a hook to send fatal messages to stdout.
func TestCrashTest_Logger(t *testing.T) {
	f := func(t *testing.T) {
		crash_LoggerFatal()
	}
	CrashTest(t, f, "fatal crash from logger")
}

func crash_NoMessage() {
	os.Exit(1)
}

func crash_ErrorMessage() {
	fmt.Println("about to crash... crashing in 3...2...1...")
	os.Exit(1)
}

func crash_ExitCode() {
	fmt.Println("crashing with a exit code != 1")
	os.Exit(2)
}

func crash_LoggerFatal() {
	// hook sends fatal messages to stdout, so they can be checked by CrashTest()
	logger, _ := HookedLogger()

	// calling Fatal() causes the process to exit
	logger.Fatal().Msg("fatal crash from logger... crashing in 3...2...1...")
}
