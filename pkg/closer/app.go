package closer

import "errors"

type appCloser struct {
	closers []closer
}

func newAppCloser(expectedClosings int) *appCloser {
	return &appCloser{make([]closer, 0, expectedClosings)}
}

func (c *appCloser) addCloser(cl closer) {
	if cl == nil {
		return
	}

	c.closers = append(c.closers, cl)
}

func (c *appCloser) closeAll() error {
	closeErrs := make([]error, 0, len(c.closers))

	for i := range c.closers {
		if c.closers[i] == nil {
			continue
		}

		if err := c.closers[i].Close(); err != nil {
			closeErrs = append(closeErrs, err)
		}
	}

	return errors.Join(closeErrs...)
}
