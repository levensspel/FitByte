package request

import "encoding/json"

type CustomString struct {
	Value     string
	IsPresent bool
	IsNull    bool
}

// This function is ran automatically on the `ctx.BodyParser()`
func (cs *CustomString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		cs.IsPresent = true
		cs.IsNull = true
		return nil
	}

	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	cs.Value = s
	cs.IsPresent = true
	cs.IsNull = false
	return nil
}

type CustomInt struct {
	Value     int
	IsPresent bool
	IsNull    bool
}

// This function is ran automatically on the `ctx.BodyParser()`
func (ci *CustomInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ci.IsPresent = true
		ci.IsNull = true
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}

	ci.Value = i
	ci.IsPresent = true
	ci.IsNull = false
	return nil
}
