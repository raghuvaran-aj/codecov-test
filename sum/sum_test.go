package sum_test

import (
	"coverageTesting/sum"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddIntegers(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should returned added number",
			args: args{10, 20},
			want: 30,
		},
		{
			name: "should returned an error",
			args: args{0, 20},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := sum.AddIntegers(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}
