package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/reactivex/rxgo/v2"
)

type User struct {
	ID       int
	Name     string
	LastName string
	Age      int
}

func main() {
	ch := make(chan rxgo.Item)

	go fetchUserIds(ch)

	obs := rxgo.FromChannel(ch)

	obsAgg := obs.Distinct(distinctUserId).Map(enrichUser).Filter(filterByAge(18)).Count()
	count := <-obsAgg.Observe()

	fmt.Println(count.V)
}

func fetchUserIds(ch chan rxgo.Item) {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- rxgo.Of(i)
		}

		for i := 20; i < 50; i++ {
			ch <- rxgo.Of(i)
		}
		close(ch)
	}()
}

func fetchUserById(id int) (User, error) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return User{
		ID:       id,
		Name:     fmt.Sprintf("name(%d)", id),
		LastName: fmt.Sprintf("lastName(%d)", id),
		Age:      r1.Intn(100),
	}, nil
}

func distinctUserId(_ context.Context, i interface{}) (interface{}, error) {
	return i, nil
}

func enrichUser(_ context.Context, i interface{}) (interface{}, error) {
	userId := i.(int)

	user, err := fetchUserById(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func filterByAge(age int) func(i interface{}) bool {
	return func(i interface{}) bool {
		user := i.(User)

		return user.Age > age
	}
}
