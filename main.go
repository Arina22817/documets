package main
import (
	"fmt"
	"documets/bin"
)

func main() {
	item := bin.NewBin("123", true, "MySecretBin")
	list := bin.NewBinList()
	list = append(list, item)
	fmt.Printf("Создан бин: %+v\n", list[0])
}
