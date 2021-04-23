package arrayrandomizer

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	//ASetting string `md:"aSetting,required"`
}

type Input struct {
	AnInput []interface{} `md:"arrayIn,required"`
}

type Output struct {
	AnOutput []interface{} `md:"arrayOut"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToArray(values["arrayIn"])
	r.AnInput = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"arrayIn": r.AnInput,
	}
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToArray(values["arrayOut"])
	o.AnOutput = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"arrayOut": o.AnOutput,
	}
}
