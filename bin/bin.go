package bin

import (
	"time"
)

type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList []Bin

func NewBin(id string, private bool, name string) Bin {
	return Bin{
		ID:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func NewBinList() BinList {
	return make(BinList, 0)
}
