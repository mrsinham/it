package main

type testSTruct struct {
	id int8
}

func (t *testSTruct) GetId() int8 {
	return t.id
}

type testInterface interface {
	GetId() int8
}

func PerformAction(t *testSTruct) {
	t.id = 8
}

func PerformActionOnCastedInterfaceIf(t testInterface) {
	if struc, ok := t.(*testSTruct); ok {
		struc.id = 8
	}
}

func PerformActionOnCastedInterfaceNoIf(t testInterface) {
	t.(*testSTruct).id = 8
}

func PerformActionOnCastedInterfaceSwitch(t testInterface) {
	switch struc := t.(type) {
	case *testSTruct:
		struc.id = 8
	}
}
