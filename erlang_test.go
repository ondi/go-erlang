//
//
//

package erlang

import (
	"testing"

	"gotest.tools/assert"
)

func Test_erlang01(t *testing.T) {
	var e float64
	e = E1(1, 1)
	assert.Assert(t, e == 0.5, e)
	e = E1(1, 2)
	assert.Assert(t, e == 0.2, e)
	e = E1(2, 2)
	assert.Assert(t, e == 0.4, e)
}
