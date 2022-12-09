package utils

func MergeSlices[T any](arrs ...[]T) []T {
	var ret []T
	for _, arr := range arrs {
		ret = append(ret, arr...)
	}

	return ret
}
