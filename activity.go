package arrayrandomizer

import (
	"math/rand"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), s, true)
	if err != nil {
		return nil, err
	}

	//ctx.Logger().Debugf("Setting: %s", s.ASetting)
	ctx.Logger().Debugf("Setting: NO SETTINGS")

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)

	ctx.Logger().Debugf("Input: %s", input.AnInput)

	if err != nil {
		return true, err
	}
	var newArray []interface{}
	var arrayLen = len(input.AnInput)
	for i := 0; i < arrayLen; i++ {
		var rnd = rand.Intn(arrayLen - i)
		println(rnd)
		ctx.Logger().Debugf("Random is: %s", rnd)
		var element = input.AnInput[rnd]
		ctx.Logger().Debugf("Movign element: %s", element)
		newArray = append(newArray, element)

		copy(input.AnInput[rnd:], input.AnInput[rnd+1:]) // Shift a[i+1:] left one index.
		//		input.AnInput[len(input.AnInput)-1] = ""             // Erase last element (write zero value).
		input.AnInput = input.AnInput[:len(input.AnInput)-1] // Truncate slice.
	}

	ctx.Logger().Debugf("Output: %s", newArray)

	output := &Output{AnOutput: newArray}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
