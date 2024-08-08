package repo

import (
	"fmt"
	"main/httpcli"
	"sync"
)

func NewUserRepo() *User {
	u := &User{
		Name: usr.Name,
		info: make(map[string]string),
	}
	for _, opt := range opts {
		opt(u)
	}
	return u
}

/*type UserInterface interface {
	setInfo(key, value string)
	initOpts(opts ...Option) func()
}*/

type UserRepo struct {
	client httpcli.CustomClient
}

func Write(originalUser *User) {
	var wg sync.WaitGroup
	threads := 10
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			user := NewUserInterface(originalUser)

			key := fmt.Sprintf("key-%d", i)
			value := fmt.Sprintf("value-%d", i)
			user.setInfo(key, value)

			fmt.Printf("%v User: %s, Info: %v\n", i, user.Name, user.info)
		}(i)
	}
	wg.Wait()
}
