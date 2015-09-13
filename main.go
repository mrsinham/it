package main

type testStruct struct {
	id int8
}

func (t *testStruct) GetId() int8 {
	return t.id
}

type testInterface interface {
	GetId() int8
}

func PerformAction(t *testStruct) {
	t.id = 8
}

func PerformActionOnCastedInterfaceIf(t testInterface) {
	if struc, ok := t.(*testStruct); ok {
		struc.id = 8
	}
}

func PerformActionOnCastedInterfaceNoIf(t testInterface) {
	t.(*testStruct).id = 8
}

func PerformActionOnCastedInterfaceSwitch(t testInterface) {
	switch struc := t.(type) {
	case *testStruct:
		struc.id = 8
	}
}
