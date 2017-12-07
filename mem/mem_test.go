package mem

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMemStats(t *testing.T) {
	stats, err := GetStats()
	require.NoError(t, err, "Getting mem stats")
	fmt.Printf("%+v", stats)
}
