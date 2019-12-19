// (c) 2019 Dapper Labs - ALL RIGHTS RESERVED

package operation

import (
	"testing"

	"github.com/dapperlabs/flow-go/crypto"
	"github.com/stretchr/testify/assert"

	"github.com/dapperlabs/flow-go/model"
	"github.com/dapperlabs/flow-go/model/flow"
)

func TestMakePrefix(t *testing.T) {

	code := byte(0x01)

	u := uint64(1337)
	expected := []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x05, 0x39}
	actual := makePrefix(code, u)

	assert.Equal(t, expected, actual)

	h := crypto.Hash{0x02, 0x03}
	expected = []byte{0x01, 0x02, 0x03}
	actual = makePrefix(code, h)

	assert.Equal(t, expected, actual)

	r := flow.Role(2)
	expected = []byte{0x01, 0x02}
	actual = makePrefix(code, r)

	assert.Equal(t, expected, actual)

	id := model.Identifier{0x05, 0x06, 0x07}
	expected = []byte{0x01,
		0x05, 0x06, 0x07, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	actual = makePrefix(code, id)

	assert.Equal(t, expected, actual)
}
