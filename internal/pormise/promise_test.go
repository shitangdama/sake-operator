package main

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	var promise = New(func(resolve func(Any), reject func(error)) {
		resolve(nil)
	})

	if promise == nil {
		t.Error("Promise is nil")
	}
}

func TestPromise_Then(t *testing.T) {
	var promise = New(func(resolve func(Any), reject func(error)) {
		time.Sleep(3 * time.Second)
		t.Log(1222222111111)
		resolve(1 + 1)
	})

	// promise.
	// 	Then(func(data Any) Any {
	// 		t.Log(4444444)
	// 		return data.(int) + 1
	// 	})

	promise.Await()
	// promise.
	// 	Then(func(data Any) Any {
	// 		return data.(int) + 1
	// 	}).
	// 	Then(func(data Any) Any {
	// 		if data.(int) != 3 {
	// 			t.Error("Result doesn't propagate")
	// 		}
	// 		return nil
	// 	}).
	// 	Catch(func(err error) error {
	// 		t.Error("Catch triggered in .Then test")
	// 		return nil
	// 	})

	// promise.Await()
}
