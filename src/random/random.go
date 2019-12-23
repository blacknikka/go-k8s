package random

import (
	"fmt"
	"math/rand"
	"time"
)

// Random ランダム
type Random interface {
	GetRandomValue() int64
	ShowRandomValue()
}

type random struct {
}

// NewRandom ランダムのインスタンス
func NewRandom() Random {
	rand.Seed(time.Now().UnixNano())
	return &random{}
}

func (r random) GetRandomValue() int64 {
	return rand.Int63()
}

func (r random) ShowRandomValue() {
	fmt.Printf("value: %v\n", rand.Int63())
}
