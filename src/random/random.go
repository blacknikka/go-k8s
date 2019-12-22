package random

import "fmt"

// Random ランダム
type Random interface {
	GetRandomValue() int64
	ShowRandomValue()
}

type random struct {
}

// NewRandom ランダムのインスタンス
func NewRandom() Random {
	return &random{}
}

func (r random) GetRandomValue() int64 {
	return 100
}

func (r random) ShowRandomValue() {
	fmt.Printf("value: %v\n", 100)
}
