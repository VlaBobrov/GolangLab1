package main

import (
	"testing"
)

type testpair struct {
	matrix [][]float32
	vector []float32
	n      int
	result []float32
}

var tests = []testpair{
	{[][]float32{{-5, 7}, {9, 8}}, []float32{29, -11}, 2, []float32{-3, 2}},
	{[][]float32{{1, 2}, {3, -1}}, []float32{8, 3}, 2, []float32{2, 3}},
}

func TestAverage(t *testing.T) {
	var A [][]float32
	var b []float32
	var E [][]float32
	var y float32
	var yes float32
	var i int = 0
	for _, pair := range tests {
		A = pair.matrix
		E = Matrix_E(E[:], pair.n)
		b = pair.vector
		E = Vector_E(pair.n, E)
		E, A, y, yes = F2(pair.n, A, E, y, yes)
		E, A, yes = F3(pair.n, A, E, yes)
		var X = F1(pair.n, b, E)
		if X[i] != pair.result[i] {
			t.Error(
				"For", pair.matrix,
				"expected", pair.result,
				"got", X,
			)
		}
		i++
	}
}
