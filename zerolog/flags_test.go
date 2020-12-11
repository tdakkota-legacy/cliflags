package zerolog

import (
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/tdakkota/cliflags/testutil"
)

func TestBuild(t *testing.T) {
	logger := Build(testutil.Parse(t, Flags, "--logging.level=trace"))

	require.Equal(t, zerolog.TraceLevel, logger.GetLevel())
}
