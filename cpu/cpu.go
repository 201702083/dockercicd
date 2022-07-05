package cpu

import (
    "fmt"
	"math/rand"
	. "dockercicd/measure"
)
func OccurCpuStress(){
  fmt.Println("\tCPU Stress test start")
  GetData()
  for i:=0 ; i < 10 ; i++ {
//	go OccurCpuStress()
	go func(){
	  big := 1000000
	  for k := 0 ; k < big ; k ++ {
		_ = rand.Intn(big)
	  }
	}()

  }
  GetData()
  fmt.Println("\tCPU Stress test end")
}

func iOccurCpuStress() {
	big := 1000000
	for i:=0 ; i < big ; i ++{
	  _ = rand.Intn(big)
	}
}
