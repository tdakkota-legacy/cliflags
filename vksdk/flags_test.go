package vksdk

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdakkota/cliflags/testutil"
)

func TestBuild(t *testing.T) {
	builder, err := Build(testutil.Parse(t, Flags, `--tokens=abcd`, `--vk.limit=10`))
	require.NoError(t, err)

	vk := builder.Complete()
	require.Equal(t, 10, vk.Limit)
}
