package windowutil

import "runtime"

type Operator interface {
	ListWindows() ([]*Window, error)
	MoveWindows(w *Window, x, y int) error
}

func NewOperator() Operator {
	var res Operator
	switch runtime.GOOS {
	case `windows`:
		res = NewWinUtil()
	case `linux`:
		res = NewLinuxUtil()
	case `darwin`:
		res = NewDarwinUtil()
	}

	return res
}
