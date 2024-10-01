package design

import (
	"context"
	"fmt"
	"time"
)

func Start() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				// do something
				fmt.Println("do something")
			}
		}
	}()
}
