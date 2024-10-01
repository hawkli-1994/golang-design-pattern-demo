package design

import "sync"

type singleton struct {}
var once sync.Once
var ins *singleton
func newSingleton() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
