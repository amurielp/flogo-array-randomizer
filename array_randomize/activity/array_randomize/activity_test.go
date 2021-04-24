package array_randomize

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/project-flogo/core/data/coerce"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestActivityRegistration(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	strs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	ints := make([]interface{}, len(strs))
	for i, s := range strs {
		ints[i] = s
	}

	tc.SetInput(arrayIn, ints)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)
	output := tc.GetOutput(arrayOut).([]interface{})

	println(output)
	for l, el := range output {
		if l != 0 {
			print(";")
		}
		print(coerce.ToInt(el))
	}
	println()
	assert.Nil(t, err)
	assert.NotEqual(t, ints, output)
	assert.Equal(t, len(ints), len(output))
}
