package main

import (
	"os"
	"runtime/pprof"
)

const (
	nbStructToAdd = 1000000
)

type testStruct struct {
	id int8
}

func (t testStruct) GetId() int8 {
	return t.id
}

type testInterface interface {
	GetId() int8
}

func main() {

	var f *os.File
	var err error
	PerformAddingInterfaceToSlice()
	f, err = os.Create("/tmp/profiling_interface")
	if err != nil {
		panic(err)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
}

func PerformAddingInterfaceToSlice() {
	var s []testInterface = make([]testInterface, 0, nbStructToAdd)
	for i := 0; i < nbStructToAdd; i++ {
		s = append(s, testStruct{2})
	}
}
