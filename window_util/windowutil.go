package windowutil

import (
	"runtime"
	"sync"
)

type Operator interface {
	ListWindows() ([]*Window, error)
	MoveWindows(w *Window, x, y int) error
}

var operator Operator
var operatorOnce sync.Once

func GetOperator() Operator {
	operatorOnce.Do(func() {
		switch runtime.GOOS {
		case `windows`:
			operator = NewWinUtil()
		case `linux`:
			operator = NewLinuxUtil()
		case `darwin`:
			operator = NewDarwinUtil()
		}
	})

	return operator
}
