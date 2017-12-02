package qr

import (
	"bytes"
	"strings"
	"testing"
)

func TestQRWritingDoesNotFail(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"QR encoding",
			args{"whatever"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := Copy(w, strings.NewReader(tt.args.text)); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
