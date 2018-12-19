package dispatchers

import "github.com/ConsenSys/hellhound-vm/hh"

func Arithmetic() []hh.Operation {
	return []hh.Operation{
		hh.NewInstruction(hh.ADD, add),
		hh.NewInstruction(hh.SUB, sub),
		hh.NewInstruction(hh.MUL, mul),
		hh.NewInstruction(hh.DIV, div),
	}
}

func add(vm hh.VM, _ *hh.Contract) error {
	x, y := vm.Stack().Pop(), vm.Stack().Peek()
	y.Add(x, y)
	return nil
}

func sub(vm hh.VM, _ *hh.Contract) error {
	x, y := vm.Stack().Pop(), vm.Stack().Peek()
	y.Sub(x, y)
	return nil
}

func mul(vm hh.VM, _ *hh.Contract) error {
	x, y := vm.Stack().Pop(), vm.Stack().Peek()
	y.Mul(x, y)
	return nil
}

func div(vm hh.VM, _ *hh.Contract) error {
	x, y := vm.Stack().Pop(), vm.Stack().Peek()
	y.Div(x, y)
	return nil
}
