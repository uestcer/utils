// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

// Package errors mirrors the standard golang "errors" module.
// Manipulate errors and provide stack trace information.
// All golang codes should using this errors package.
package errors

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

// Default value if the error code is not defined.
// The error code is DefaultErrCode if Error was created by New, Newf, Wrap, Wrapf.
const (
	DefaultErrCode = -1
)

// Error interface exposes additional information about the error.
type Error interface {

	// This returns the error message without the stack trace.
	Message() string

	// This returns the stack trace without the error message.
	Stack() string

	// This returns the stack trace's context.
	Context() string

	// This returns the error code.
	// If the error code is not pre-defined, return is DefaultErrCode.
	Code() int

	// This returns the wrapped error. Nil if not wrap another error.
	Inner() error

	// Implements the built-in error interface.
	Error() string
}

// Base standard struct for interface 'Error'.
type baseError struct {
	message string
	stack   string
	context string
	code    int
	inner   error
}

// This returns the error string without stack trace information.
func Message(err interface{}) string {
	switch e := err.(type) {
	case Error:
		err := Error(e)
		ret := []string{}
		for err != nil {
			ret = append(ret, err.Message())
			innerErr := err.Inner()

			if innerErr == nil {
				break
			}

			var ok bool
			err, ok = innerErr.(Error)
			if !ok {
				ret = append(ret, innerErr.Error())
				break
			}
		}
		return strings.Join(ret, " ")
	case runtime.Error:
		return runtime.Error(e).Error()
	default:
		return "Passed a non-error to Message"
	}
}

// This returns the error message without the stack trace.
func (e *baseError) Message() string {
	return e.message
}

// This returns the stack trace without the error message.
func (e *baseError) Stack() string {
	return e.stack
}

// This returns the stack trace's context.
func (e *baseError) Context() string {
	return e.context
}

// This set the error code.
func (e *baseError) SetCode(code int) {
	e.code = code
}

// This returns the error code.
func (e *baseError) Code() int {
	return e.code
}

// This returns the wrapped error, if there is one.
func (e *baseError) Inner() error {
	return e.inner
}

// This returns a string with all available error information,
// including inner errors that are wrapped by this errors.
func (e *baseError) Error() string {
	return DefaultError(e)
}

// A default implementation of the Error method of the error interface.
func DefaultError(e Error) string {

	errLines := []string{"ERROR:"}
	var origStack string
	fillErrorInfo(e, &errLines, &origStack)

	errLines = append(errLines, "")
	errLines = append(errLines, "ORIGINAL STACK TRACE:")
	errLines = append(errLines, origStack)

	return strings.Join(errLines, "\n")
}

// Fills errLines with all error messages, and origStack with the inner-most stack.
func fillErrorInfo(err error, errLines *[]string, origStack *string) {
	if err == nil {
		return
	}

	e, ok := err.(Error)
	if ok {
		*errLines = append(*errLines, e.Message())
		*origStack = e.Stack()
		fillErrorInfo(e.Inner(), errLines, origStack)
	} else {
		*errLines = append(*errLines, err.Error())
	}
}

// This returns a new baseError initialized with the given message and
// the current stack trace.
func New(msg string) Error {
	stack, context := StackTrace()
	return &baseError{
		message: msg,
		stack:   stack,
		context: context,
		code:    DefaultErrCode,
	}
}

// This returns a new baseError initialized with the given message, error code and
// the current stack trace.
func NewByCode(code int, msg string) Error {
	stack, context := StackTrace()
	return &baseError{
		message: msg,
		stack:   stack,
		context: context,
		code:    code,
	}
}

// Same as New, but with fmt.Printf-style parameters.
func Newf(format string, args ...interface{}) Error {
	stack, context := StackTrace()
	return &baseError{
		message: fmt.Sprintf(format, args...),
		stack:   stack,
		context: context,
		code:    DefaultErrCode,
	}
}

// Same as NewByCode, but with fmt.Printf-style parameters.
func NewfByCode(code int, format string, args ...interface{}) Error {
	stack, context := StackTrace()
	return &baseError{
		message: fmt.Sprintf(format, args...),
		stack:   stack,
		context: context,
		code:    code,
	}
}

// Wraps another error in a new baseError.
func Wrap(err error, msg string) Error {
	stack, context := StackTrace()
	return &baseError{
		message: msg,
		stack:   stack,
		context: context,
		inner:   err,
		code:    DefaultErrCode,
	}
}

// Wraps another error in a new baseError with error code information.
func WrapByCode(code int, err error, msg string) Error {
	stack, context := StackTrace()
	return &baseError{
		message: msg,
		stack:   stack,
		context: context,
		inner:   err,
		code:    code,
	}
}

// Same as Wrap, but with fmt.Printf-style parameters.
func Wrapf(err error, format string, args ...interface{}) Error {
	stack, context := StackTrace()
	return &baseError{
		message: fmt.Sprintf(format, args...),
		stack:   stack,
		context: context,
		inner:   err,
		code:    DefaultErrCode,
	}
}

// Same as WrapByCode, but with fmt.Printf-style parameters.
func WrapfByCode(code int, err error, format string, args ...interface{}) Error {
	stack, context := StackTrace()
	return &baseError{
		message: fmt.Sprintf(format, args...),
		stack:   stack,
		context: context,
		inner:   err,
		code:    code,
	}
}

// Returns a copy of the error with the stack trace field populated and any
// other shared initialization; skips 'skip' levels of the stack trace.
// NOTE: This panics on any error.
func stackTrace(skip int) (current, context string) {

	buf := make([]byte, 128)
	for {
		n := runtime.Stack(buf, false)
		if n < len(buf) {
			buf = buf[:n]
			break
		}
		buf = make([]byte, len(buf)*2)
	}

	indexNewline := func(b []byte, start int) int {
		if start >= len(b) {
			return len(b)
		}
		searchBuf := b[start:]
		index := bytes.IndexByte(searchBuf, '\n')
		if index == -1 {
			return len(b)
		} else {
			return (start + index)
		}
	}

	var strippedBuf bytes.Buffer
	index := indexNewline(buf, 0)
	if index != -1 {
		strippedBuf.Write(buf[:index])
	}

	for i := 0; i < skip; i++ {
		index = indexNewline(buf, index+1)
		index = indexNewline(buf, index+1)
	}

	isDone := false
	startIndex := index
	lastIndex := index
	for !isDone {
		index = indexNewline(buf, index+1)
		if (index - lastIndex) <= 1 {
			isDone = true
		} else {
			lastIndex = index
		}
	}

	strippedBuf.Write(buf[startIndex:index])
	return strippedBuf.String(), string(buf[index:])
}

// This returns the current stack trace string.
func StackTrace() (current, context string) {
	return stackTrace(3)
}
