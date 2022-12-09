package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMergeSlices(t *testing.T) {
	var s1 = []string{"1", "2"}
	var s2 = []string{"3"}
	var s3 = make([]string, 0)
	var s4 []string
	var s5 []string = nil

	require.Len(t, MergeSlices(s1, s2), 3)
	require.Equal(t, []string{"1", "2", "3"}, MergeSlices(s1, s2))
	require.Len(t, MergeSlices(s1, s3), 2)
	require.Len(t, MergeSlices(s3), 0)
	require.Len(t, MergeSlices(s2, s4), 1)
	require.Len(t, MergeSlices(s1, s2, s5), 3)
	require.Len(t, MergeSlices(s1, s2, nil), 3)
}
