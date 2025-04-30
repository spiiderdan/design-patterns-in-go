package singleton

import(
	"sync"
)

var once sync.Once

type Counter struct{
	value int
}

var instance *Counter

func GetInstance() *Counter{
	once.Do(func() {
		instance = &Counter{}
	})
	return instance
}

func (c *Counter) Add(){
	c.value++
}

func (c *Counter) Get() int{
	return c.value
}