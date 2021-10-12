package main

import (
	"context"
	"fmt"
	"sync"
)

// Any is a substitute for interface{}
type Any = interface{}

type response struct {
	res Any
	err error
}

// A Promise is a proxy for a value not necessarily known when
// the promise is created. It allows you to associate handlers
// with an asynchronous action's eventual success value or failure reason.
// This lets asynchronous methods return values like synchronous methods:
// instead of immediately returning the final value, the asynchronous method
// returns a promise to supply the value at some point in the future.
type Promise struct {
	pending bool

	ctx context.Context

	// A function that is passed with the arguments resolve and reject.
	// The executor function is executed immediately by the Promise implementation,
	// passing resolve and reject functions (the executor is called
	// before the Promise constructor even returns the created object).
	// The resolve and reject functions, when called, resolve or reject
	// the promise, respectively. The executor normally initiates some
	// asynchronous work, and then, once that completes, either calls the
	// resolve function to resolve the promise or else rejects it if
	// an error or panic occurred.
	executor func(resolve func(Any), reject func(error))

	result chan *response
	// Stores the result passed to resolve()
	// result Any

	// Stores the error passed to reject()
	err error

	// Mutex protects against data race conditions.
	mutex sync.Mutex

	// WaitGroup allows to block until all callbacks are executed.
	wg sync.WaitGroup
}

// New instantiates and returns a pointer to a new Promise.
func New(executor func(resolve func(Any), reject func(error))) *Promise {

	// 这里还要解决一下超时

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var promise = &Promise{
		ctx: ctx,
		// cancel:   cancel,
		pending:  true,
		executor: executor,
		result:   make(chan *response, 1),
		err:      nil,
		mutex:    sync.Mutex{},
		// wg:       sync.WaitGroup{},
	}
	// promise.wg.Add(1)

	defer close(promise.result)
	fmt.Println(7777)
	go func(ctx context.Context) {
		// defer promise.handlePanic()
		fmt.Println(888888)
		promise.executor(promise.resolve, promise.reject)
	}(promise.ctx)

	return promise
}

// 一个自函数 主要用用于处理内置的resolver方法
func (promise *Promise) resolve(resolution Any) {
	promise.mutex.Lock()

	if !promise.pending {
		promise.mutex.Unlock()
		return
	}

	switch result := resolution.(type) {
	case *Promise:
		flattenedResult, err := result.Await()
		fmt.Println(flattenedResult)
		if err != nil {
			promise.mutex.Unlock()
			promise.reject(err)
			return
		}
		// promise.result = flattenedResult

		// promise.result <- &response{
		// 	result: flattenedResult,
		// }
	default:
		// promise.result <- &response{
		// 	result: result,
		// }
		// promise.result = result
	}
	promise.pending = false

	promise.wg.Done()
	promise.mutex.Unlock()
}

// 内部的reject方法
func (promise *Promise) reject(err error) {
	promise.mutex.Lock()
	defer promise.mutex.Unlock()

	if !promise.pending {
		return
	}

	promise.err = err
	promise.pending = false

	promise.ctx.Done()
}

// func (promise *Promise) handlePanic() {
// 	e := recover()

// 	if e != nil {
// 		switch err := e.(type) {
// 		case nil:
// 			promise.reject(fmt.Errorf("panic recovery with nil error"))
// 		case error:
// 			promise.reject(fmt.Errorf("panic recovery with error: %s", err.Error()))
// 		default:
// 			promise.reject(fmt.Errorf("panic recovery with unknown error: %s", fmt.Sprint(err)))
// 		}
// 	}
// }

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

// Await is a blocking function that waits for all callbacks to be executed.
// Returns value and error.
// Call on an already resolved Promise to get its result and error
func (promise *Promise) Await() (Any, error) {
	// promise.wg.Wait()

	for {
		select {
		case <-promise.ctx.Done():
			fmt.Println(33333)
			fmt.Println(22222222)
			return promise.result, promise.err
		case r := <-promise.result:
			if r.err == nil {
				// 	out <- r
				fmt.Println(11111)
			}
			return promise.result, promise.err
		}
	}

	return promise.result, promise.err
}

// 这是两个对外的函数

// // Resolve returns a Promise that has been resolved with a given value.
// func Resolve(resolution Any) *Promise {
// 	return New(func(resolve func(Any), reject func(error)) {
// 		resolve(resolution)
// 	})
// }

// // Reject returns a Promise that has been rejected with a given error.
// func Reject(err error) *Promise {
// 	return New(func(resolve func(Any), reject func(error)) {
// 		reject(err)
// 	})
// }
