package closer

import "errors"

type appCloser struct {
	closers []closer
}

func (c *appCloser) closeAll() error {
	closeErrs := make([]error, 0, len(c.closers))

	for i := range c.closers {
		if err := c.closers[i].Close(); err != nil {
			closeErrs = append(closeErrs, err)
		}
	}

	return errors.Join(closeErrs...)
}
