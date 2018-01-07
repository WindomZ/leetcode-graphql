package leetcodegraphql

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestBaseQuestion_Do1(t *testing.T) {
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

func TestBaseQuestion_Do2(t *testing.T) {
	q := new(BaseQuestion)

	assert.NoError(t, q.Do("1"))

	code, err := q.GetCodeDefinition("golang")
	assert.NoError(t, err)
	assert.NotEmpty(t, code)
	t.Logf("code: %s", code)

	env, err := q.GetEnvInfo("golang")
	assert.NoError(t, err)
	assert.NotEmpty(t, env)
	t.Logf("env: %s", env)
}
