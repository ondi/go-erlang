//
//
//

package erlang

import (
	"testing"

	"gotest.tools/assert"
)

func Test_E1_01(t *testing.T) {
	var e float64
	e = E1(1, 1)
	assert.Assert(t, e == 0.5, e)
	e = E1(1, 2)
	assert.Assert(t, e == 0.2, e)
	e = E1(2, 2)
	assert.Assert(t, e == 0.4, e)
}

func Test_V1_01(t *testing.T) {
	var v float64
	var err error
	v, err = V1(50, 0.01)
	assert.NilError(t, err)
	assert.Assert(t, v == 64, v)
	v, err = V1(1, 0.5)
	assert.NilError(t, err)
	assert.Assert(t, v == 1, v)
	// v, err = V1(1000000000, 0.0001)
	// assert.NilError(t, err)
	// assert.Assert(t, v == 999909251, v)
}

func Test_V2_01(t *testing.T) {
	var v float64
	var err error
	v, err = V2(50, 0.01)
	assert.NilError(t, err)
	assert.Assert(t, v == 68, v)
	v, err = V2(1, 0.5)
	assert.NilError(t, err)
	assert.Assert(t, v == 2, v)
}
