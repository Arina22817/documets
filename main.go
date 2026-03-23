package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList []Bin

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
}

func NewBinList() BinList {
	return make(BinList, 0)
}

func main() {
	item := NewBin("123", true, "MySecretBin")
	list := NewBinList()
	list = append(list, item)
	fmt.Printf("Создан бин: %+v\n", list[0])
}
