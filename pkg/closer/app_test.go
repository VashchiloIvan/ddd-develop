package closer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_appCloser_closeAll(t *testing.T) {
	t.Parallel()

	err1 := errors.New("err1")
	err2 := errors.New("err2")

	type fields struct {
		closers []closer
	}
	tests := []struct {
		name     string
		fields   fields
		wantErrs []error
	}{
		{
			name: "close with some errs",
			fields: fields{
				closers: []closer{
					newFuncCloser(func() error {
						return nil
					}),
					newFuncCloser(func() error {
						return err1
					}),
					newFuncCloser(func() error {
						return nil
					}),
					newFuncCloser(func() error {
						return err2
					}),
				},
			},
			wantErrs: []error{err1, err2},
		},
		{
			name: "close without errs",
			fields: fields{
				closers: []closer{
					newFuncCloser(func() error {
						return nil
					}),
					newFuncCloser(func() error {
						return nil
					}),
				},
			},
			wantErrs: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &appCloser{
				closers: tt.fields.closers,
			}

			err := c.closeAll()
			if err == nil {
				assert.True(t, len(tt.wantErrs) == 0)
			} else {
				assert.False(t, len(tt.wantErrs) == 0)

				for _, expectedErr := range tt.wantErrs {
					assert.ErrorIs(t, err, expectedErr)
				}
			}
		})
	}
}

func Test_appCloser_addCloser(t *testing.T) {
	t.Parallel()

	type fields struct {
		closers []closer
	}
	type args struct {
		cl closer
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantClosersCount int
	}{
		{
			name: "add nil closer",
			fields: fields{
				closers: nil,
			},
			args: args{
				cl: nil,
			},
		},
		{
			name: "add real closer",
			fields: fields{
				closers: nil,
			},
			args: args{
				cl: newFuncCloser(func() error {
					return nil
				}),
			},
			wantClosersCount: 1,
		},
		{
			name: "add at list of closers",
			fields: fields{
				closers: []closer{
					newFuncCloser(func() error {
						return nil
					}), newFuncCloser(func() error {
						return nil
					}), newFuncCloser(func() error {
						return nil
					}),
				},
			},
			args: args{
				cl: newFuncCloser(func() error {
					return nil
				}),
			},
			wantClosersCount: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &appCloser{
				closers: tt.fields.closers,
			}
			c.addCloser(tt.args.cl)

			assert.Equal(t, tt.wantClosersCount, len(c.closers))
		})
	}
}
