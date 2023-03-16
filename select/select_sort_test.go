package _select

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var num = 10

func TestSelectSort(t *testing.T) {
	data := testData()
	sorted := make([][]int, 0)
	for _, item := range data {
		var a = make([]int, 0)
		a = append(a, item...)
		sorted = append(sorted, Sort(a))
	}

	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
		fmt.Println(sorted[i])
		fmt.Println()
	}

}

func testData() [][]int {
	data := make([][]int, 0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < num; i++ {
		item := make([]int, 0)
		for j := 0; j < 10; j++ {
			item = append(item, rand.Intn(100))
		}
		data = append(data, item)
	}

	return data
}
