package main

import (
	"fmt"
	"net/http"
	Cpu "dockercicd/cpu"
	Mem "dockercicd/mem"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("[log] someone connected")
		w.Write([]byte("I wanna break pods\nGo to /memStress & /cpuStress"))
	})
	http.HandleFunc("/cpuStress",func(w http.ResponseWriter, req *http.Request){
		Cpu.OccurCpuStress()
		w.Write([]byte("cpu is burning"))
	})
	http.HandleFunc("/memStress", func(w http.ResponseWriter, req *http.Request){
		Mem.OccurMemStress()
		w.Write([]byte("mem is melting"))
	})

	http.ListenAndServe(":5000", nil)
}
