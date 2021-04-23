package main

import (
	"math/rand"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// Constants
const (
	arrayIn  = "arrayIn"
	arrayOut = "arrayOut"
)

var activityLog = logger.GetLogger("array_randomizer")

type ArrRandActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &ArrRandActivity{metadata: metadata}
}

func (a *ArrRandActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *ArrRandActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing array randomizer")

	if context.GetInput(arrayIn) == nil {
		return false, activity.NewError("arrayIn is mandatory", "ARR_RAND-1000", nil)
	}
	input := context.GetInput(arrayIn).([]interface{})

	activityLog.Info("Input: %s", input)

	if err != nil {
		return true, err
	}
	var newArray []interface{}
	var arrayLen = len(input)
	for i := 0; i < arrayLen; i++ {
		var rnd = rand.Intn(arrayLen - i)
		//println(rnd)
		activityLog.Info("Random is: %s", rnd)
		var element = input[rnd]
		activityLog.Info("Movign element: %s", element)
		newArray = append(newArray, element)

		copy(input[rnd:], input[rnd+1:]) // Shift a[i+1:] left one index.
		//		input[len(input)-1] = ""             // Erase last element (write zero value).
		input = input[:len(input)-1] // Truncate slice.
	}
	activityLog.Info("Output: %s", newArray)

	context.SetOutput(arrayOut, newArray)

	return true, nil
}
