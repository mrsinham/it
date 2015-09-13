package main

import "testing"

func BenchmarkPerformActionOnStruct(b *testing.B) {
	b.ReportAllocs()
	t := testStruct{2}
	for i := 0; i < b.N; i++ {
		PerformAction(&t)
	}
}

func BenchmarkPerformActionOnCastedInterfaceIf(b *testing.B) {
	b.ReportAllocs()
	t := testStruct{2}
	for i := 0; i < b.N; i++ {
		PerformActionOnCastedInterfaceIf(&t)
	}
}

func BenchmarkPerformActionOnCastedInterfaceNoIf(b *testing.B) {
	b.ReportAllocs()
	t := testStruct{2}
	for i := 0; i < b.N; i++ {
		PerformActionOnCastedInterfaceNoIf(&t)
	}
}

func BenchmarkPerformActionOnCastedInterfaceSwitch(b *testing.B) {
	b.ReportAllocs()
	t := testStruct{2}
	for i := 0; i < b.N; i++ {
		PerformActionOnCastedInterfaceSwitch(&t)
	}
}
