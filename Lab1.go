package main

import "fmt"

func Matrix_Inicialization(A [][]float32, n int) [][]float32 {
	fmt.Println("Enter Matrix :")
	for index := 0; index < n; index++ {
		tmp := make([]float32, n)
		for j := 0; j < n; j++ {
			fmt.Scanf("%f", &tmp[j])
		}
		A = append(A, tmp)
	}
	return A[:]
}
func Matrix_E(E [][]float32, n int) [][]float32 {
	for index := 0; index < n; index++ {
		tmp := make([]float32, n)
		for j := 0; j < n; j++ {
			tmp[j] = 0
		}
		E = append(E, tmp)
	}
	return E[:]
}
func Vector_Inicialization(b []float32, n int) []float32 {
	fmt.Println("Enter Vector :")
	tmp := make([]float32, n)
	for j := 0; j < n; j++ {
		fmt.Scanf("%f", &tmp[j])
	}
	b = append(tmp)
	return b[:]
}
func Vector_X(n int) []float32 {
	var X []float32
	tmp := make([]float32, n)
	for j := 0; j < n; j++ {
		tmp[j] = 0
	}
	X = append(tmp)
	return X[:]
}

func F1(n int, b []float32, E [][]float32) []float32 {
	var X []float32 = Vector_X(n)
	for i := 0; i < n; i++ {
		var S float32 = 0
		for j := 0; j < n; j++ {
			S = S + E[i][j]*b[j]
			X[i] = S
		}
		fmt.Printf("%0.3f \t", X[i])
	}
	return X[:]
}
func Vector_E(n int, E [][]float32) [][]float32 {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				E[i][j] = 1

			} else {
				E[i][j] = 0
			}
		}
	}
	return E[:]
}

func F2(n int, A [][]float32, E [][]float32, y float32, yes float32) ([][]float32, [][]float32, float32, float32) {

	for k := 0; k < n; k++ {
		if A[k][k] == 0 {
			y = 0
		} else {
			y = 1 / A[k][k]
		}
		for j := 0; j < n; j++ {
			A[k][j] = A[k][j] * y
			E[k][j] = E[k][j] * y
		}

		for i := k + 1; i < n; i++ {
			yes = A[i][k]
			for z := 0; z < n; z++ {
				A[i][z] = A[i][z] - A[k][z]*yes
				E[i][z] = E[i][z] - E[k][z]*yes
			}
		}
	}
	return E[:], A[:], y, yes
}
func F3(n int, A [][]float32, E [][]float32, yes float32) ([][]float32, [][]float32, float32) {
	for k := n - 1; k >= 0; k-- {
		for i := k - 1; i >= 0; i-- {
			yes = A[i][k]
			for z := n - 1; z >= 0; z-- {
				A[i][z] = A[i][z] - A[k][z]*yes
				E[i][z] = E[i][z] - E[k][z]*yes
			}
		}
	}
	return E[:], A[:], yes
}

/*****-5 7 9 8 29 -11*********/
func main() {
	var A [][]float32
	var b []float32
	fmt.Println("Please Enter matrix size from 2 to 4")
	var n int
	fmt.Scanf("%d", &n)
	var E [][]float32
	var y float32
	var yes float32
	A = Matrix_Inicialization(A[:], n)
	E = Matrix_E(E[:], n)
	b = Vector_Inicialization(b[:], n)
	E = Vector_E(n, E)
	E, A, y, yes = F2(n, A, E, y, yes)
	E, A, yes = F3(n, A, E, yes)
	F1(n, b, E)
}
