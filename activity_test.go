package arrayrandomizer

import (
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	strs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ints := make([]interface{}, len(strs))
	for i, s := range strs {
		ints[i] = s
	}

	input := &Input{AnInput: ints}
	err := tc.SetInputObject(input)
	assert.Nil(t, err)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)
	output := &Output{}
	err = tc.GetOutputObject(output)

	println(output.AnOutput)
	for l, el := range output.AnOutput {
		if l != 0 {
			print(";")
		}
		print(coerce.ToInt(el))
	}
	println()
	assert.Nil(t, err)
	assert.NotEqual(t, ints, output.AnOutput)
	assert.Equal(t, len(ints), len(output.AnOutput))
}
