package measure
import (
	"time"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"math"
	"fmt"
	"strconv"
	"log"
)

func getCpuUsage()int{
	  percent,err:=cpu.Percent(time.Second,false)
	  if err!=nil{
		  log.Fatal(err)
	  }
	return int(math.Ceil(percent[0]))
}

func getMemoryUsage()int{
	  memory,err:=mem.VirtualMemory()
	  if err!=nil{
	  	log.Fatal(err)
	  }
	return int(math.Ceil(memory.UsedPercent))
}

func GetData(){
	cpuData:="Cpu: "+strconv.Itoa(getCpuUsage())+"% "
	memoryData:="Mem: "+strconv.Itoa(getMemoryUsage())+"% "
	fmt.Println(cpuData+memoryData)
}
