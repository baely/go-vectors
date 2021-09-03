package vector

import "errors"

type Vector struct {
	values []int
}

func Create(values ...int) Vector {
	newVector := Vector{values: values}
	return newVector
}

func Add(v1, v2 Vector) (Vector, error) {
	if len(v1.values) != len(v2.values) {
		return Vector{}, errors.New("can't add vectors of different cardinality")
	}

	newValues := make([]int, len(v1.values))

	for i := range newValues {
		newValues[i] = v1.values[i] + v2.values[i]
	}

	newVector := Vector{values: newValues}

	return newVector, nil
}

func Subtract(v1, v2 Vector) (Vector, error) {
	if len(v1.values) != len(v2.values) {
		return Vector{}, errors.New("can't subtract vectors of different cardinality")
	}

	newValues := make([]int, len(v1.values))

	for i := range newValues {
		newValues[i] = v1.values[i] - v2.values[i]
	}

	newVector := Vector{values: newValues}

	return newVector, nil
}

func Multiply(v Vector, k int) Vector {
	newValues := make([]int, len(v.values))

	for i := range newValues {
		newValues[i] = v.values[i] * k
	}

	newVector := Vector{values: newValues}
	return newVector
}

func DotProduct(v1, v2 Vector) (int, error) {
	if len(v1.values) != len(v2.values) {
		return 0, errors.New("can't calculate dot product of vectors of different cardinality")
	}

	newValue := 0

	for i := 0; i < len(v1.values); i++ {
		newValue += v1.values[i] * v2.values[i]
	}

	return newValue, nil
}
