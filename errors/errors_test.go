// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package errors

import (
	"fmt"
	"strings"
	"testing"
)

func TestStackTrace(t *testing.T) {
	const testMsg = "test error"
	er := New(testMsg)
	e := er.(*baseError)

	if e.message != testMsg {
		t.Error("error message %s != expected %s", e.message, testMsg)
	}

	if strings.Index(e.stack, "errors/errors.go") != -1 {
		t.Error("stack trace generation code should not be in the error stack trace")
	}

	if strings.Index(e.stack, "TestStackTrace") == -1 {
		t.Error("stack trace must have test code in it")
	}

	var err error = e
	_ = err
}

func TestWrappedError(t *testing.T) {
	const (
		innerMsg  = "I am inner error"
		middleMsg = "I am the middle error"
		outerMsg  = "I am the mighty outer error"
	)

	inner := fmt.Errorf(innerMsg)
	middle := Wrap(inner, middleMsg)
	outer := Wrap(middle, outerMsg)
	errorStr := outer.Error()

	if strings.Index(errorStr, innerMsg) == -1 {
		t.Errorf("couldn't find inner error message in:\n%s", errorStr)
	}

	if strings.Index(errorStr, middleMsg) == -1 {
		t.Errorf("couldn't find middle error message in:\n%s", errorStr)
	}

	if strings.Index(errorStr, outerMsg) == -1 {
		t.Errorf("couldn't find outer error message in:\n%s", errorStr)
	}
}
