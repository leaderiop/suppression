package main

import (
	"fmt"

	"github.com/leaderiop/suppression/domain/models"
)

func main() {

	tree := models.NewTree()
	data := []float32{1, 9, 2, 8, 3, 7, 4, 6, 5}
	tree.SetParentValue(data[0])
	for _, v := range data[1:] {
		tree.AppendValue(v)
	}
	fmt.Println(tree.HasValue(1), tree.HasValue(15), tree.HasValue(8))
}
