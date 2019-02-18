package main

import "testing"
import "reflect"

func TestSquare(t *testing.T) {
	sq := square{}
	if sq.sideLength != 0 {
		t.Errorf("Uninizialized Square Should have length of 0")
	}

	ft := reflect.TypeOf(sq.sideLength).Kind()
	if ft != reflect.Float64 {
		t.Errorf("Square's sideLength should be type of float64")
	}

	sq = square {
		sideLength: 1,
	}

	if sq.getArea() != 1.0 {
		t.Errorf("Square Area does not equal 1.0")
	}
}

func TestTriangle(t *testing.T) {
	tr := triangle{}
	ft := reflect.TypeOf(tr.base).Kind()
	if ft != reflect.Float64 {
		t.Errorf("%v %v should be type of float64", "triangle's", "base")
	}

	ft = reflect.TypeOf(tr.height).Kind()
	if ft != reflect.Float64 {
		t.Errorf("%v %v should be type of float64", "triangle's", "height")
	}

	tr = triangle{
		base : 1.0,
		height : 1.0,
	}

	if tr.getArea() != 0.5 {
		t.Errorf("Triangle get area should equal 0.5")
	}

	tr = triangle{
		base : 2.0,
		height : 1.0,
	}

	if tr.getArea() != 1.0 {
		t.Errorf("Triangle get area should equal 1.0")
	}
}