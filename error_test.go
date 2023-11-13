package aselerror

import (
	"fmt"
	"testing"
)

var (
	ErrorValidate = &AselError{
		cause: "Parameter verification failed",
		code:  100401,
		err:   nil,
		level: 1,
	}
)

func TestNew(t *testing.T) {
	err := ErrorValidate

	fmt.Printf("error: %s\n", err)

	fmt.Printf("hide error cause: %s\n", err.HideDetails())

	var testError error

	testError = err

	fmt.Printf("test another err: %s\n", testError)

	err.ShowDetails()

	fmt.Printf("test another err show: %s\n", testError)

	// === RUN   TestNew
	// error: Parameter verification failed
	// hide error cause: 100401
	// test another err: 100401
	// test another err show: Parameter verification failed
	// --- PASS: TestNew (0.00s)
	// PASS
}
