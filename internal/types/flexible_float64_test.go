package types_test

import (
	"encoding/json"
	"testing"

	"github.com/amadejkastelic/mercator-api/internal/types"
)

func TestFlexibleFloat64_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{
			name:    "null value",
			input:   "null",
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "float64 value",
			input:   "123.456",
			want:    123.456,
			wantErr: false,
		},
		{
			name:    "integer value",
			input:   "123",
			want:    123.0,
			wantErr: false,
		},
		{
			name:    "string float",
			input:   `"123.456"`,
			want:    123.456,
			wantErr: false,
		},
		{
			name:    "string integer",
			input:   `"123"`,
			want:    123.0,
			wantErr: false,
		},
		{
			name:    "empty string",
			input:   `""`,
			want:    0.0,
			wantErr: false,
		},
		{
			name:    "invalid string",
			input:   `"invalid"`,
			want:    0.0,
			wantErr: true,
		},
		{
			name:    "boolean value",
			input:   "true",
			want:    0.0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f types.FlexibleFloat64
			err := json.Unmarshal([]byte(tt.input), &f)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && float64(f) != tt.want {
				t.Errorf("UnmarshalJSON() got = %v, want %v", float64(f), tt.want)
			}
		})
	}
}

func TestFlexibleFloat64_String(t *testing.T) {
	tests := []struct {
		name  string
		value types.FlexibleFloat64
		want  string
	}{
		{
			name:  "positive float",
			value: types.FlexibleFloat64(123.456),
			want:  "123.456",
		},
		{
			name:  "negative float",
			value: types.FlexibleFloat64(-123.456),
			want:  "-123.456",
		},
		{
			name:  "integer value",
			value: types.FlexibleFloat64(123),
			want:  "123",
		},
		{
			name:  "zero value",
			value: types.FlexibleFloat64(0),
			want:  "0",
		},
		{
			name:  "large value",
			value: types.FlexibleFloat64(1e10),
			want:  "10000000000",
		},
		{
			name:  "small value",
			value: types.FlexibleFloat64(1e-10),
			want:  "0.0000000001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.value.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFlexibleFloat64_MarshalUnmarshal(t *testing.T) {
	tests := []struct {
		name  string
		value float64
	}{
		{
			name:  "positive float",
			value: 123.456,
		},
		{
			name:  "negative float",
			value: -123.456,
		},
		{
			name:  "integer value",
			value: 123,
		},
		{
			name:  "zero value",
			value: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a struct with our flexible type
			type TestStruct struct {
				Value types.FlexibleFloat64 `json:"value"`
			}

			// Create and marshal the original
			original := TestStruct{Value: types.FlexibleFloat64(tt.value)}
			data, err := json.Marshal(original)
			if err != nil {
				t.Fatalf("Failed to marshal: %v", err)
			}

			// Unmarshal back and compare
			var decoded TestStruct
			if err := json.Unmarshal(data, &decoded); err != nil {
				t.Fatalf("Failed to unmarshal: %v", err)
			}

			if float64(decoded.Value) != tt.value {
				t.Errorf("Round trip failed: got %v, want %v", float64(decoded.Value), tt.value)
			}
		})
	}
}

// Test for a struct with multiple FlexibleFloat64 fields
func TestFlexibleFloat64_StructUnmarshal(t *testing.T) {
	type TestStruct struct {
		IntValue    types.FlexibleFloat64 `json:"int_value"`
		FloatValue  types.FlexibleFloat64 `json:"float_value"`
		StringValue types.FlexibleFloat64 `json:"string_value"`
		EmptyValue  types.FlexibleFloat64 `json:"empty_value"`
	}

	jsonData := `{
		"int_value": 123,
		"float_value": 456.789,
		"string_value": "789.123",
		"empty_value": ""
	}`

	var result TestStruct
	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal struct: %v", err)
	}

	if float64(result.IntValue) != 123.0 {
		t.Errorf("IntValue = %v, want %v", float64(result.IntValue), 123.0)
	}

	if float64(result.FloatValue) != 456.789 {
		t.Errorf("FloatValue = %v, want %v", float64(result.FloatValue), 456.789)
	}

	if float64(result.StringValue) != 789.123 {
		t.Errorf("StringValue = %v, want %v", float64(result.StringValue), 789.123)
	}

	if float64(result.EmptyValue) != 0.0 {
		t.Errorf("EmptyValue = %v, want %v", float64(result.EmptyValue), 0.0)
	}
}
