package main

type testStruct struct {
	id int16
}

func (t *testStruct) GetId() int16 {
	return t.id
}

type testInterface interface {
	GetId() int16
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
