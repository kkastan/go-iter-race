package main

import (
	"fmt"
	"sync"

	"github.com/Hearst-DD/spackler"
)

func doPrint(val string) {
	fmt.Printf("%s\n", val)
}

func main() {
	caddy := spackler.New(true)
	var wg sync.WaitGroup
	tokens := []string{"Charlie", "Mac", "Dennis", "Dee", "Frank"}

	for _, token := range tokens {
		my_token := token
		wg.Add(1)
		caddy.Go(func(c *spackler.Caddy) {
			defer wg.Done()
			doPrint(my_token)
		})
	}

	wg.Wait()
}
