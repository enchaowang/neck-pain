package windowutil

type LinuxUtil struct {
}

func NewLinuxUtil() *LinuxUtil {
	// todo init of win
	return nil
}

func (w *LinuxUtil) ListWindows() ([]*Window, error) {
	panic("not implemented") // TODO: Implement
}

func (w *LinuxUtil) MoveWindows(win *Window, x int, y int) error {
	panic("not implemented") // TODO: Implement
}
