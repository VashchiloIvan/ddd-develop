package closer

type funcCloser struct {
	f func() error
}

func newFuncCloser(f func() error) *funcCloser {
	return &funcCloser{f: f}
}

func (c *funcCloser) Close() error {
	return c.f()
}
