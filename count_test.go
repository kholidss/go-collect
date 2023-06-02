package collect

import (
	"errors"
	"testing"
)

func TestCount(t *testing.T) {
	tests := []struct {
		Value          interface{}
		ExpectedResult int
		WantError      error
	}{
		{
			Value:          "Hello",
			ExpectedResult: 5,
			WantError:      nil,
		},

		{
			Value:          12345678,
			ExpectedResult: 8,
			WantError:      nil,
		},
		{
			Value:          int8(53),
			ExpectedResult: 2,
			WantError:      nil,
		},
		{
			Value:          int16(5311),
			ExpectedResult: 4,
			WantError:      nil,
		},
		{
			Value:          int32(531144551),
			ExpectedResult: 9,
			WantError:      nil,
		},
		{
			Value:          int64(5311445511231234521),
			ExpectedResult: 19,
			WantError:      nil,
		},
		{
			Value: struct {
				name string
				age  int
			}{
				name: "John Doe",
				age:  30,
			},
			ExpectedResult: 2,
			WantError:      nil,
		},
		{
			Value: map[string]int{
				"one":   1,
				"two":   2,
				"three": 3,
				"four":  4,
			},
			ExpectedResult: 4,
			WantError:      nil,
		},
		{
			Value: map[int]interface{}{
				0: 1,
				1: true,
				2: int32(1234),
				3: 4,
				4: "four",
			},
			ExpectedResult: 5,
			WantError:      nil,
		},
		{
			Value:          []string{"john", "doe", "greg", "wilmen", "george"},
			ExpectedResult: 5,
			WantError:      nil,
		},
		{
			Value:          [3]int{100, 12, 8},
			ExpectedResult: 3,
			WantError:      nil,
		},
		{
			Value:          "",
			ExpectedResult: 0,
			WantError:      nil,
		},
		{
			Value:          true,
			ExpectedResult: 0,
			WantError: errWrongDataType{
				process:    "count",
				actualType: "bool",
			},
		},
	}

	for testNumb, test := range tests {

		res, err := Count(test.Value)
		if !errors.Is(err, test.WantError) {
			t.Errorf("For test #%d, got error: %v", testNumb, err)
		}

		if res != test.ExpectedResult {
			t.Errorf("For test #%d, not match with expected value. Expected: %d, got Actual: %d", testNumb, test.ExpectedResult, res)
			continue
		}
	}
}

func TestCountBy(t *testing.T) {
	tests := []struct {
		Value          []interface{}
		ExpectedResult map[interface{}]int
	}{
		{
			Value: []interface{}{1, 2, 3, 4},
			ExpectedResult: map[interface{}]int{
				1: 1,
				2: 1,
				3: 1,
				4: 1,
			},
		},
		{
			Value: []interface{}{2, "test", 10, "test", 2},
			ExpectedResult: map[interface{}]int{
				2:      2,
				"test": 2,
				10:     1,
			},
		},
		{
			Value: []interface{}{"hello", 10, "10", 14.5},
			ExpectedResult: map[interface{}]int{
				"hello": 1,
				10:      1,
				"10":    1,
				14.5:    1,
			},
		},
		{
			Value: []interface{}{true, false, true, 900},
			ExpectedResult: map[interface{}]int{
				true:  2,
				false: 1,
				900:   1,
			},
		},
	}

	for testNumb, test := range tests {

		res := CountBy(test.Value)

		if len(res) != len(test.ExpectedResult) {
			t.Errorf("For test #%d, length not match with expected value. Expected: %d, got Actual: %d", testNumb, test.ExpectedResult, res)
			continue
		}

		for k, valRes := range res {
			valTest, ok := test.ExpectedResult[k]
			if !ok || valRes != valTest {
				t.Errorf("For test #%d, value not match with expected value. Expected: %d, got Actual: %d", testNumb, test.ExpectedResult, res)
			}
		}
	}
}
