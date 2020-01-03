package domain

import "fmt"

type RandomItem struct {
	Name   string
	Random string
}

func (r *RandomItem) ToString() string {
	return fmt.Sprintf("%+v", r)
}
