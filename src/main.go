package main

import (
	"fmt"
	"time"

	"github.com/blacknikka/go-k8s/random"
)

func main() {
	rnd := random.NewRandom()
	doRandom(rnd)
}

func doRandom(rnd random.Random) {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			randomizedInt := rnd.GetRandomValue()
			fmt.Printf("value: %v\n", randomizedInt)
		}
	}
	t.Stop()
}
