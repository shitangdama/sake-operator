package main

import (
	"fmt"
	"sync"
	"time"
)

type Any = interface{}

type response struct {
	res Any
	err error
}

type Promise struct {
	pending bool

	final chan struct{}

	executor func(resolve func(Any), reject func(error))

	// result chan *response

	result Any

	err error

	mutex sync.Mutex

	wg sync.WaitGroup
}

// https://blog.csdn.net/inthat/article/details/106917358
// 这个地方可能要穿进去一个
func New(executor func(resolve func(Any), reject func(error))) *Promise {

	var promise = &Promise{

		pending:  true,
		executor: executor,
		result:   make(chan *response, 1),
		final:    make(chan struct{}, 1),
		err:      nil,
		mutex:    sync.Mutex{},
		// wg:       sync.WaitGroup{},
	}
	// promise.wg.Add(1)

	// defer close(promise.result)
	fmt.Println(7777)
	go func() {

		fmt.Println(888888)
		promise.executor(promise.resolve, promise.reject)
	}()

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
		promise.result = flattenedResult

		// promise.result <- &response{
		// 	result: flattenedResult,
		// }
	default:
		// promise.result <- &response{
		// 	result: result,
		// }
		promise.result = result
	}
	promise.pending = false

	promise.final <- struct{}{}
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

	promise.final <- struct{}{}
}

func (promise *Promise) Cancel() {
	promise.final <- struct{}{}
}

// https://github.com/fanliao/go-promise/blob/1890db352a72f9e6c6219c20111355cddc795297/future.go#L137

func (promise *Promise) Await() (Any, error) {
	// promise.wg.Wait()
	fmt.Println(123123)
	for {
		select {
		case <-promise.final:
			return promise.result, promise.err
		}
	}

	return promise.result, promise.err
}

// then就是新建一个promise

// 这里要考虑下
// 实现then
// 实现 OnSuccess， OnFailure， OnComplete， OnCancel
// addCallback
// All, Race, Reduce

func main() {
	var promise = New(func(resolve func(Any), reject func(error)) {
		time.Sleep(10 * time.Second)
		fmt.Println(1222222111111)
		resolve(1 + 1)
	})

	// promise.
	// 	Then(func(data Any) Any
	// 		t.Log(4444444)
	// 		return data.(int) + 1
	// 	})

	r, err := promise.Await()

	fmt.Println(err)
	fmt.Println(r)
}
