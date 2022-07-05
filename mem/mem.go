package mem

import (
    "fmt"
	. "dockercicd/measure"
)
func OccurMemStress() {
    fmt.Println("\tMem stress test start")
    GetData()
    var overall [][]int
    for i := 0; i<40; i++ {
        a := make([]int, 0, 999999999)
        overall = append(overall, a)
		if (i % 10 == 0) {
		  GetData()
		}
	}
    GetData()
    fmt.Println("\tMem stress test end")
    overall = nil // clear
}
