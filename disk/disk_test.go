package disk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiskStats(t *testing.T) {
	stats, err := GetStats()
	require.NoError(t, err, "Getting disk stats")
	for _, s := range stats {
		fmt.Printf("%+v\n", s)
	}
}
