package windowutil

type DarwinUtil struct {
}

func NewDarwinUtil() *DarwinUtil {
	// todo init of win
	return nil
}

func (w *DarwinUtil) ListWindows() ([]*Window, error) {
	panic("not implemented") // TODO: Implement
}

func (w *DarwinUtil) MoveWindows(win *Window, x int, y int) error {
	panic("not implemented") // TODO: Implement
}
