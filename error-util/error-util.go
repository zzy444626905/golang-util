package errutil

import (
	"github.com/pkg/errors"
	"runtime/debug"
)

func PanicMessage(message string) {
	Panic(errors.New(message))
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicWithMessage(err error,msg string) {
	if err != nil {
		panic(errors.WithMessage(err, msg))
	}
}

func Ignore(i ...interface{}) {

}
