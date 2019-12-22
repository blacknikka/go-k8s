package main

import (
	"time"

	"github.com/blacknikka/go-k8s/random"
)

var rnd random.Random

func main() {
	rnd = random.NewRandom()
	doRandom()
}

func doRandom() {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-t.C:
			rnd.ShowRandomValue()
		}
	}
	t.Stop()
}
