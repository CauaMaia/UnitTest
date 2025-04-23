package pitoco

import "testing"

func TestInsertGrades(t *testing.T) {
	tests := []struct {
		grades        []int
		expected      bool
		expectedError string
		expectedAvg   float64
	}{
		{grades: []int{5, 10, 0}, expected: true, expectedError: "", expectedAvg: 5.0},
		{grades: []int{5, -1, 10}, expected: false, expectedError: "invalid grade found", expectedAvg: 0.0},
		{grades: []int{11, 5, 7}, expected: false, expectedError: "invalid grade found", expectedAvg: 0.0},
	}

	for _, test := range tests {
		result := InsertGrades(test.grades)
		if result != test.expected {
			t.Errorf("InsertGrades(%v) = %v; want %v", test.grades, result, test.expected)
		}

		average := calculateAverage(test.grades)
		if isValidAverage(test.grades) {
			if average != test.expectedAvg {
				t.Errorf("Expected average %.2f but got %.2f for grades %v", test.expectedAvg, average, test.grades)
			}
			t.Logf("The average of grades %v is %.2f and is valid", test.grades, average)
		} else {
			if test.expectedError != "" {
				t.Logf("Expected error: %s for grades %v", test.expectedError, test.grades)
			}
			t.Logf("The average of grades %v is %.2f and is invalid", test.grades, average)
		}
	}
}

func calculateAverage(grades []int) float64 {
	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}

func isValidAverage(grades []int) bool {
	for _, grade := range grades {
		if grade < 0 || grade > 10 {
			return false
		}
	}
	return true
}
