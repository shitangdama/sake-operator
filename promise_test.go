package main

import "testing"

func TestNew(t *testing.T) {
	var promise = New(func(resolve func(Any), reject func(error)) {
		resolve(nil)
	})

	if promise == nil {
		t.Error("Promise is nil")
	}
}
