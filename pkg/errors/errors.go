package errors

import (
	"fmt"
	"os"
)

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func ThrowError(msg string) {
	err := New(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}
func ThrowErrorMsg(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", msg)
		os.Exit(1)
	}
}
func Waring(msg string) {
	err := New(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Waring: %s\n", err.Error())
	}
}
func WaringErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func ReturnError(err error) bool {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		return false
	}
	return true
}
