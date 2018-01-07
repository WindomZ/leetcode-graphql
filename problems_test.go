package leetcodegraphql

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestProblems_Do(t *testing.T) {
	p := new(Problems)
	assert.NoError(t, p.Do())

	s1 := p.StatStatus("1")
	assert.NotEmpty(t, s1)
	assert.Equal(t, "Two Sum", s1.Stat.QuestionTitle)

	s2 := p.StatStatus("Two Sum")
	assert.NotEmpty(t, s2)
	assert.Equal(t, "Two Sum", s2.Stat.QuestionTitle)

	s3 := p.StatStatus("two-sum")
	assert.NotEmpty(t, s3)
	assert.Equal(t, "Two Sum", s3.Stat.QuestionTitle)
}
