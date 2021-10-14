package main

import (
	"errors"
	"math/rand"
	"sync"
	"unsafe"
)

// Any is a substitute for interface{}
type Any = interface{}

// https://github.com/fanliao/go-promise/blob/1890db352a72f9e6c6219c20111355cddc795297/future.go#L137

//Future provides a read-only view of promise,
//the value is set by using Resolve, Reject and Cancel methods of related Promise
type Promise struct {
	ID    int //Id can be used as identity of Future
	final chan struct{}

	executor func(resolve func(Any), reject func(error))

	result interface{} // The Promise's data.
	err    error       // The error status.

	//val point to futureVal that stores status of future
	//if need to change the status of future, must copy a new futureVal and modify it,
	//then use CAS to put the pointer of new futureVal
	// 这个是不是状态
	val unsafe.Pointer

	valMap sync.Map
}

// New instantiates and returns a pointer to a new Promise.
func New(executor func(resolve func(Any), reject func(error))) *Promise {

	promise := &Promise{
		ID:       rand.Int(),
		executor: executor,
		final:    make(chan struct{}),
		// unsafe.Pointer(val),
	}
	go func() {

		// defer promise.handlePanic()
		promise.executor(promise.resolve, promise.reject)
	}()

	return promise
}

func (promise *Promise) resolve(resolution Any) {
	switch result := resolution.(type) {
	case *Promise:
		flattenedResult, err := result.Await()
		if err != nil {
			promise.reject(err)
			return
		}
		promise.result = flattenedResult
	default:
		promise.result = result
	}
	close(promise.final)
}

func (promise *Promise) reject(err error) {
	promise.err = err
	close(promise.final)
}

// Await is xxx
func (promise *Promise) Await() (val interface{}, err error) {
	<-promise.final
	// return getFutureReturnVal(this.loadResult())

	// Block until at least one of these conditions is satisfied. If both are,
	// "select" will choose one pseudo-randomly.
	// select {
	// case <-p.final:
	return promise.result, promise.err
	// }
}

// Cancel xxxxxx
func (promise *Promise) Cancel() {
	close(promise.final)
	promise.err = errors.New("promise cancel")
}

// Catch Appends a rejection handler to the promise,
// and returns a new promise resolving to the return value of the handler.
func (promise *Promise) Catch(rejection func(err error) error) *Promise {
	return New(func(resolve func(Any), reject func(error)) {
		result, err := promise.Await()
		if err != nil {
			reject(rejection(err))
			return
		}
		resolve(result)
	})
}

// Then appends fulfillment and rejection handlers to the promise,
// and returns a new promise resolving to the return value of the called handler.
func (promise *Promise) Then(fulfillment func(data Any) Any) *Promise {
	return New(func(resolve func(Any), reject func(error)) {
		result, err := promise.Await()
		if err != nil {
			reject(err)
			return
		}
		resolve(fulfillment(result))
	})
}

//Race ... resolves to the very first promise, rejects if none of the promises resolves
func Race(promises []*Promise) *Promise {

}
