package application_test_test

import (
	"backend-academy_2024_project_1-go-progirira/internal/application"
	"testing"
)

func TestInitLives(t *testing.T) {
	tests := map[string]struct {
		input         int
		expectedValue int
		expectedErr   error
	}{
		"positive number": {
			input:         7,
			expectedValue: 7,
			expectedErr:   nil,
		},
		"number is zero": {
			input:         0,
			expectedValue: 0,
			expectedErr:   application.NewErrNonPositiveAttempts(0),
		},
		"negative number": {
			input:         -5,
			expectedValue: 0,
			expectedErr:   application.NewErrNonPositiveAttempts(-5),
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lives, err := application.InitLives(tc.input)

			if tc.expectedErr != nil {
				if err == nil || err.Error() != tc.expectedErr.Error() {
					t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
				}

				return
			}

			if err != nil {
				t.Errorf("expected error: %v, got: %v", tc.expectedErr, err)
			}

			if int(*lives) != tc.expectedValue {
				t.Errorf("got value: %v, expected value: %v", int(*lives), tc.expectedValue)
			}
		})
	}
}

func TestDecreaseLives(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected int
	}{
		"positive number": {
			input:    7,
			expected: 6,
		},
		"zero": {
			input:    0,
			expected: -1,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lives := application.Lives(tc.input)
			lives.DecreaseLives()

			if int(lives) != tc.expected {
				t.Errorf("got: %v, expected: %v", int(lives), tc.expected)
			}
		})
	}
}

func TestIsHasLives(t *testing.T) {
	tests := map[string]struct {
		input    int
		expected bool
	}{
		"has lives": {
			input:    7,
			expected: true,
		},
		"has no lives": {
			input:    0,
			expected: false,
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lives := application.Lives(tc.input)
			got := lives.IsHasLives()

			if got != tc.expected {
				t.Errorf("got: %v, expected: %v", got, tc.expected)
			}
		})
	}
}
