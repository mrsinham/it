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
	PerformAddingStructureToSlice()
	f, err = os.Create("/tmp/profiling_struct")
	if err != nil {
		panic(err)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
}

func PerformAddingStructureToSlice() {
	var s []testStruct = make([]testStruct, 0, nbStructToAdd)
	for i := 0; i < nbStructToAdd; i++ {
		s = append(s, testStruct{2})
	}
}
