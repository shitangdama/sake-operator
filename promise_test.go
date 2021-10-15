package main

import (
	"testing"
	"time"
)

func TestNewP(t *testing.T) {
	var promise = New(func(promise *Promise) {
		promise.resolve(nil)
	})

	if promise == nil {
		t.Error("Promise is nil")
	}
}

func TestNewPromise(t *testing.T) {
	var promise = New(func(promise *Promise) {
		time.Sleep(10 * time.Second)
		t.Log(1222222111111)
		promise.resolve(1 + 1)
	})

	result, err := promise.Await()
	t.Log(result)
	t.Log(err)
}

func TestNewPromiseThen(t *testing.T) {
	var promise = New(func(promise *Promise) {

		time.Sleep(10 * time.Second)
		t.Log(1222222111111)
		promise.resolve(1)
	}).Then(func(data Any) Any {
		t.Log(4444444)
		return data.(int) + 1
	}).
		Then(func(data Any) Any {
			t.Log(55555555)
			if data.(int) != 3 {
				t.Error("Result doesn't propagate")
			}
			return data
		}).
		Catch(func(err error) error {
			t.Log(66666)
			t.Error("Catch triggered in .Then test")
			return nil
		})

	result, err := promise.Await()
	t.Log(result)
	t.Log(err)
}

// func TestPromiseThenNested(t *testing.T) {
// 	var promise = New(func(resolve func(Any), reject func(error)) {
// 		resolve(New(func(res func(Any), rej func(error)) {
// 			res("Hello, World")
// 		}))
// 	})

// 	promise.
// 		Then(func(data Any) Any {
// 			if data.(string) != "Hello, World" {
// 				t.Error("Resolved promise doesn't flatten")
// 			}
// 			return nil
// 		}).
// 		Catch(func(err error) error {
// 			t.Error("Catch triggered in .Then test")
// 			return nil
// 		})

// 	result, err := promise.Await()
// 	t.Log(result)
// 	t.Log(err)
// }

// // https://github.com/chebyrash/promise/blob/master/promise_test.go

// func TestPromiseCatchNested(t *testing.T) {
// 	var promise = New(func(resolve func(Any), reject func(error)) {
// 		resolve(New(func(res func(Any), rej func(error)) {
// 			rej(errors.New("nested fail"))
// 		}))
// 	})

// 	promise.
// 		Then(func(data Any) Any {
// 			t.Error("Then triggered in .Catch test")
// 			return nil
// 		}).
// 		Catch(func(err error) error {
// 			if err.Error() != "nested fail" {
// 				t.Error("Rejected promise doesn't flatten")
// 			}
// 			return nil
// 		})

// 	result, err := promise.Await()
// 	t.Log(result)
// 	t.Log(err)
// }
