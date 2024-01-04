package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_convertTimeFormat(t *testing.T) {
	type args struct {
		inputTime string
	}

	tests := []struct {
		name        string
		args        args
		want        string
		expectedErr bool
	}{
		{
			name:        "AM time test case",
			args:        args{"12:45:30AM"},
			want:        "00:45:30",
			expectedErr: false,
		},
		{
			name:        "PM time test case",
			args:        args{"09:59:30PM"},
			want:        "21:59:30",
			expectedErr: false,
		},

		{
			name:        "Invalid format test case",
			args:        args{"5:59PM"},
			want:        "",
			expectedErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertTimeFormat(tt.args.inputTime)
			if tt.expectedErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}

}

func BenchmarkConvertTimeFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		convertTimeFormat("12:01:00AM")
	}
}
