package vector

import "errors"

type Vector struct {
	values []int
	len int
}

func Create(values ...int) Vector {
	newVector := Vector{values: values, len: len(values)}
	return newVector
}

func Add(v1 Vector, v2 Vector) (Vector, error) {
	if v1.len != v2.len {
		return Vector{}, errors.New("can't add vectors of different cardinality")
	}

	newValues := make([]int, v1.len)

	for i := 0; i < v1.len; i++ {
		newValues[i] = v1.values[i] + v2.values[i]
	}

	newVector := Vector{values: newValues, len: v1.len}

	return newVector, nil
}

func Subtract(v1 Vector, v2 Vector) (Vector, error) {
	if v1.len != v2.len {
		return Vector{}, errors.New("can't subtract vectors of different cardinality")
	}

	newValues := make([]int, v1.len)

	for i := 0; i < v1.len; i++ {
		newValues[i] = v1.values[i] - v2.values[i]
	}

	newVector := Vector{values: newValues, len: v1.len}

	return newVector, nil
}

func DotProduct(v1 Vector, v2 Vector) (int, error) {
	if v1.len != v2.len {
		return 0, errors.New("can't calculate dot product of vectors of cardinality")
	}

	newValue := 0

	for i := 0; i < v1.len; i++ {
		newValue += v1.values[i] * v2.values[i]
	}

	return newValue, nil
}
