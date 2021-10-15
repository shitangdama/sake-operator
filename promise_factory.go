package main

// //All ... resolves a Promise when all promises passed are resolved,
// func All(promises []*Promise) *Promise {

// 	return New(func(resolve func(Any), reject func(error))) {

// 		var w sync.WaitGroup

// 		// 这里要不要传入promise类直接进来
// 		// data := make([]interface{}, len(promises))

// 		for i, promise := range promises {
// 			index := i // because go catches current i value not the ones that was encountered when loop was at this loop state

// 			if promise == nil {
// 				data[index] = nil
// 				w.Done()
// 				continue
// 			}

// 		}

// 		w.Wait()

// 	})
// }

// //Race ... resolves to the very first promise, rejects if none of the promises resolves
func Race(promises []*Promise) *Promise {

	return nil
}
