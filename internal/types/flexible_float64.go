package types

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type FlexibleFloat64 float64

func (f *FlexibleFloat64) UnmarshalJSON(data []byte) error {
	// Try if float64
	var fl float64
	if err := json.Unmarshal(data, &fl); err == nil {
		*f = FlexibleFloat64(fl)
		return nil
	}

	// Try if int
	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*f = FlexibleFloat64(i)
		return nil
	}

	// Lastly try if string
	var str string
	if err := json.Unmarshal(data, &str); err == nil {
		if str == "" {
			*f = FlexibleFloat64(0)
			return nil
		}
		if fl, err := strconv.ParseFloat(str, 64); err == nil {
			*f = FlexibleFloat64(fl)
			return nil
		}
	}

	return fmt.Errorf("failed to unmarshal FlexibleFloat64: %s", string(data))
}

func (f FlexibleFloat64) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}
