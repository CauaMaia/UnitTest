package pitoco

// insert grade
func InsertGrades(grades []int) bool {
	for _, grade := range grades {
		if grade < 0 || grade > 10 {
			return false
		}
	}
	return true
}

// calculate average of grades
func AverageGrades(grades []int) float64 {
	if len(grades) == 0 {
		return 0
	}

	sum := 0
	for _, grade := range grades {
		sum += grade
	}
	return float64(sum) / float64(len(grades))
}
