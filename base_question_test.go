package leetcodegraphql

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestBaseQuestion_Do(t *testing.T) {
	q := new(BaseQuestion)

	assert.NoError(t, q.Do("two-sum"))

	code, err := q.GetCodeDefinition("golang")
	assert.NoError(t, err)
	assert.NotEmpty(t, code)
	t.Logf("code: %s", code)

	env, err := q.GetEnvInfo("golang")
	assert.NoError(t, err)
	assert.NotEmpty(t, env)
	t.Logf("env: %s", env)
}
