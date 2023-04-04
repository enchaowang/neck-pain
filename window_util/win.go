package windowutil

type WinUtil struct {
}

func NewWinUtil() *WinUtil {
	// todo init of win
	return nil
}

func (w *WinUtil) ListWindows() ([]*Window, error) {
	panic("not implemented") // TODO: Implement
}

func (w *WinUtil) MoveWindows(win *Window, x int, y int) error {
	panic("not implemented") // TODO: Implement
}
