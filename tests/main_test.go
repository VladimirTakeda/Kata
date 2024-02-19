package tests

import "testing"
import utils "Kata/utils"

func TestEvaluateWithPanics(t *testing.T) {
	// Тестовый случай с арабским числом 0
	t.Run("One number", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("1")
	})

	// Тестовый случай с арабским числом 11
	t.Run("Mixed systems", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("1 + II")
	})

	t.Run("More then two params", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("1 + 2 + 3")
	})

	t.Run("Out of limits first param Arab", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("11 + 2")
	})

	t.Run("Out of limits second param Arab", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("1 + 22")
	})

	t.Run("Missed space 1", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("1 +2")
	})

	t.Run("Missed space 2", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("1+ 2")
	})

	t.Run("Missed space 3", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("1+2")
	})

	t.Run("Negative Roman result", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("I - II")
	})

	t.Run("Float number", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		utils.Calculate("1.1 - 2")
	})
}

func TestEvaluateCorrect(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"1 + 1\n", "2"},
		{"1 - 10\n", "-9"},
		{"10 / 3\n", "3"},
		{"10 * 10\n", "100"},
		{"1 / 1\n", "1"},
		{"7 / 2\n", "3"},
		{"6 / 2\n", "3"},
		{"4 / 10\n", "0"},
		{"1 - 1\n", "0"},

		{"I + I\n", "II"},
		{"X / III\n", "III"},
		{"X * X\n", "C"},
		{"I / I\n", "I"},
		{"VII / II\n", "III"},
		{"VI / II\n", "III"},
		{"V * X\n", "L"},
	}

	for _, testCase := range testCases {
		result := utils.Calculate(testCase.input)
		if result != testCase.expected {
			t.Errorf("For input %s, expected %s, got %s", testCase.input, testCase.expected, result)
		}
	}
}
