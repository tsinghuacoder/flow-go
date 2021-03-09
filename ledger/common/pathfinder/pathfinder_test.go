package pathfinder_test

import (
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/crypto/hash"
	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common/pathfinder"
)

// Test_KeyToPathV0 tests key to path for V0
func Test_KeyToPathV0(t *testing.T) {

	kp1 := ledger.KeyPartFixture(1, "key part 1")
	kp2 := ledger.KeyPartFixture(22, "key part 2")
	k := ledger.NewKey([]ledger.KeyPart{kp1, kp2})

	path, err := pathfinder.KeyToPath(k, 0)
	require.NoError(t, err)

	// compute expected value
	h := sha256.New()
	_, err = h.Write([]byte("key part 1"))
	require.NoError(t, err)
	_, err = h.Write([]byte("key part 2"))
	require.NoError(t, err)
	expected := ledger.Path(h.Sum(nil))

	require.True(t, path.Equals(expected))
}

func Test_KeyToPathV1(t *testing.T) {

	kp1 := ledger.KeyPartFixture(1, "key part 1")
	kp2 := ledger.KeyPartFixture(22, "key part 2")
	k := ledger.NewKey([]ledger.KeyPart{kp1, kp2})

	path, err := pathfinder.KeyToPath(k, 1)
	require.NoError(t, err)

	// compute expected value
	hasher := hash.NewSHA3_256()
	_, err = hasher.Write([]byte("/1/key part 1/22/key part 2"))
	require.NoError(t, err)

	expected := ledger.Path(hasher.SumHash())
	require.True(t, path.Equals(expected))
}
