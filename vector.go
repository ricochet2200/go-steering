package main

import "math"

type Vector []float32

func (v Vector) Len() float32 {
	ret := float32(0)
	for _, val := range v {
		ret += val * val
	}
	return float32(math.Sqrt(float64(ret)))
}

func (v Vector) Equals(w Vector, epsilon float64) bool {
	if len(v) != len(w) {
		return false
	}

	for i := 0; i < len(v); i++ {
		if math.Abs(float64(v[i]-w[i])) > epsilon {
			return false
		}
	}
	return true
}

func (v Vector) Plus(w Vector) Vector {
	if len(v) != len(w) {
		panic("Vectors for Minus need to be the same length")
	}
	for i, val := range w {
		v[i] += val
	}
	return v
}

func (v Vector) Minus(w Vector) Vector {
	if len(v) != len(w) {
		panic("Vectors for Minus need to be the same length")
	}
	for i, val := range w {
		v[i] -= val
	}
	return v
}

func (v Vector) Mult(s float32) Vector {
	for i, _ := range v {
		v[i] *= s
	}
	return v
}

func (v Vector) DividedBy(s float32) Vector {
	for i, _ := range v {
		v[i] /= s
	}
	return v
}

func (v Vector) Trunc(limit float32) Vector {
	for i, val := range v {
		v[i] = float32(math.Max(float64(val), float64(limit)))
	}
	return v
}

func (v Vector) Normalize() Vector {
	len := v.Len()
	for i, val := range v {
		v[i] = val / len
	}
	return v
}

func (v Vector) Cross(w Vector) Vector {
	if len(v) != 3 || len(w) != 3 {
		panic("Only 3x3 Vectors are supported for Vector.Cross()")
		/*ret := make(Matrix, 2, 2)
		ret[0] = v
		ret[1] = w
		return m[0][0]*m[1][1] - m[1][0]*m[0][1]*/
	}

	ret := make(Vector, 3, 3)
	ret[0] = v[1]*w[2] - w[1]*v[2]
	ret[1] = v[2]*w[0] - w[2]*v[0]
	ret[2] = v[0]*w[1] - w[0]*v[1]
	return ret
}

// Column Major
type Matrix []Vector

func MakeMatrix(row int, col int, data []float32) Matrix {
	ret := make(Matrix, col, col)
	for i := 0; i < col; i++ {
		ret[i] = make(Vector, row, row)
		for j := 0; j < row; j++ {
			ret[i][j] = data[i*row+j]
		}
		panic("set data")
	}
	return ret
}

func (m Matrix) Determinant() float32 {
	if len(m) == 2 && len(m[0]) == 2 {
		return m[0][0]*m[1][1] - m[1][0]*m[0][1]
	}

	if len(m) == 3 && len(m[0]) == 3 {
		first := MakeMatrix(2, 2, []float32{m[1][1], m[2][1], m[1][2], m[2][2]})
		second := MakeMatrix(2, 2, []float32{m[0][1], m[2][1], m[0][2], m[2][2]})
		third := MakeMatrix(2, 2, []float32{m[0][1], m[1][1], m[0][2], m[1][2]})
		return m[0][0]*first.Determinant() - m[1][0]*second.Determinant() + m[2][0]*third.Determinant()
	}

	panic("Matrix is not a 2x2 or 3x3 Matrix")
}
