package aselerror

import (
	"fmt"
	"testing"
)

var (
	TestError = New(
		"parameter verification failed",
		WithCode(100401),
		WithLevel(1),
	)
)

func TestNew(t *testing.T) {
	t.Logf("test err { Error() string } method: %s", TestError)

	result := func() (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = New(
					"panic error",
					WithCode(100401),
					WithLevel(8),
					WithCause(fmt.Errorf("%v", r)),
				)
			}
		}()

		panic("test panice error")
	}

	err := result()
	if e, ok := err.(AselError); ok {
		if e.GetLevel() > 3 {
			t.Logf("warning: %s", e)
		} else {
			t.Logf("normal: %s", e)
		}
	} else {
		t.Errorf("unexpected mistakes")
	}
}
