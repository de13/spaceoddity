package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"time"
)

var cpu, mem bool
var second int
var startTime time.Time

func init() {
	startTime = time.Now()
	flag.BoolVar(&cpu, "cpu", false, "Simulate a CPU intensive application.")
	flag.BoolVar(&mem, "mem", false, "Simulate a memory leak.")
	flag.IntVar(&second, "second", 300, "Time (in seconds) for the simulation.")
	flag.IntVar(&second, "s", 300, "Time (in seconds) for the simulation.")
}

func main() {
	flag.Parse()
	if cpu && mem {
		fmt.Println("cpu and mem are mutually exclusive.")
		os.Exit(1)
	}
	if cpu {
		cpuIntensive(second)
	}
	if mem {
		memoryLeak(second)
	}
	for {
		time.Sleep(3600 * time.Second)
	}
}

func cpuIntensive(duration int) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	n := runtime.NumCPU()
	runtime.GOMAXPROCS(n)

	for i := 0; i < n; i++ {
		go func() {
			for uptime() < time.Duration(duration)*time.Second {
				fmt.Fprintf(f, ".")
			}
		}()
	}
}

func memoryLeak(duration int) {
	dd()
	var fileSlice []string
	file, err := ioutil.ReadFile("10m_file")
	if err != nil {
		log.Fatal("File not found.")
	}
	for i := 0; i < duration; i++ {
		fileSlice = append(fileSlice, string(file))
		time.Sleep(1 * time.Second)
	}
	fileSlice = nil
	err = os.Remove("10m_file")
	if err != nil {
		log.Fatal("Failed to remove file.")
	}
}

func uptime() time.Duration {
	return time.Since(startTime)
}

func dd() {
	size := int64(10 * 1024 * 1024)
	fd, err := os.Create("10m_file")
	defer fd.Close()
	if err != nil {
		log.Fatal("Failed to create file.")
	}
	_, err = fd.Seek(size-1, 0)
	if err != nil {
		log.Fatal("Failed to seek file.")
	}
	_, err = fd.Write([]byte{0})
	if err != nil {
		log.Fatal("Failed to write file.")
	}
}
