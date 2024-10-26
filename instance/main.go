package instance

import "sync"

func CreateInstance[T any](Init func() *T) func() *T {
	var once sync.Once
	var instance *T
	return func() *T {
		once.Do(func() {
			instance = Init()
		})
		return instance
	}
}

// func New(paran NewParam, requsetParam RequestParam) *Instance {
// 	return &Instance

// }

// type NewParam struct {
// }
// type RequestParam struct {
// }
// type Instance struct {
// }
