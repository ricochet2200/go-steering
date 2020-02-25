package main

import (
	"math"
	"testing"
)

func TestLen(t *testing.T) {
	v := Vector{3, 4}
	if v.Len() != 5 {
		t.Errorf("The Length of Vector{3,4} should be 5")
	}

	v2 := Vector{8, 0}
	if v2.Len() != 8 {
		t.Errorf("The Length of Vector{8,0} should be 8")
	}

	v3 := Vector{1, 2, 3}
	if math.Abs(float64(v3.Len()-3.742)) > .001 {
		t.Errorf("The Length of Vector{1,2,3} should be ~0.3742")
	}
}

func TestEquals(t *testing.T) {
	v := Vector{1, 2, 3}
	w := Vector{1, 5, 7}
	if v.Equals(w, .01) {
		t.Errorf("%g should not equal %g", v, w)
	}

	v = Vector{1, 2, 3}
	w = Vector{1, 2, 3.0000001}
	if !v.Equals(w, .01) {
		t.Errorf("%g should equal %g", v, w)
	}

	w = Vector{1, 2, 3.1}
	if v.Equals(w, .01) {
		t.Errorf("%g epsilon is smaller than the difference in %g", v, w)
	}

}

func TestCross(t *testing.T) {
	v := Vector{1, 2, 3}
	w := Vector{1, 5, 7}
	a := Vector{-1, -4, 3}

	if !v.Cross(w).Equals(a, 0) {
		t.Errorf("Cross is not correct")
	}
}

func TestNormalize(t *testing.T) {
	v := Vector{3, 1, 2}
	a := Vector{.802, .267, .534}
	if v.Normalize().Equals(a, .0001) {
		t.Errorf("Vector Normalization doesn't work")
	}
	if v[0] != 3 {
		t.Errorf("Vector Normalization changed the vector to %f", v[0])
	}

}
