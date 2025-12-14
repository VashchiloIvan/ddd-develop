package closer

import (
	"errors"
	"testing"
)

func Test_funcCloser_Close(t *testing.T) {
	t.Parallel()

	type fields struct {
		f func() error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "successful close result",
			fields: fields{
				f: func() error {
					return nil
				},
			},
		},
		{
			name: "fail close result",
			fields: fields{
				f: func() error {
					return errors.New("failed to close")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			c := &funcCloser{
				f: tt.fields.f,
			}
			if err := c.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
