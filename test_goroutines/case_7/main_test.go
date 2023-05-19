package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Calc(t *testing.T) {
	t.Parallel()

	t.Run("test", func(t *testing.T) {
		t.Parallel()
		result := 3

		res := CalcSum(1, 2)

		assert.Equal(t, result, res)
	})
}

func TestCalcSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{1, 2},
			want: 3,
		},
		{
			name: "test2",
			args: args{3, 2},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CalcSum(tt.args.x, tt.args.y), "CalcSum(%v, %v)", tt.args.x, tt.args.y)
		})
	}
}
