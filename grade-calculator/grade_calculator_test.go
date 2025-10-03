package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 40, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 40, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment String() = %q, want %q", Assignment.String(), "assignment")
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam String() = %q, want %q", Exam.String(), "exam")
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay String() = %q, want %q", Essay.String(), "essay")
	}
}

func TestComputeAverageEmpty(t *testing.T) {
	avg := computeAverage([]Grade{})
	if avg != 0 {
		t.Fatalf("avg=%d, want 0 for empty slice", avg)
	}
}

func TestGradeCalculation(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a1", 100, Assignment)
	gc.AddGrade("e1", 80, Exam)
	gc.AddGrade("s1", 60, Essay)

	got := gc.calculateNumericalGrade()
	if got != 87 {
		t.Fatalf("numerical=%d, want 87", got)
	}
}

func TestGetFinalGradeBoundaries(t *testing.T) {
	cases := []struct {
		score int
		want  string
	}{
		{90, "A"},
		{80, "B"},
		{70, "C"},
		{60, "D"},
	}

	for _, tc := range cases {
		gc := NewGradeCalculator()
		gc.AddGrade("a", tc.score, Assignment)
		gc.AddGrade("e", tc.score, Exam)
		gc.AddGrade("s", tc.score, Essay)

		if got := gc.GetFinalGrade(); got != tc.want {
			t.Fatalf("score=%d => got %q, want %q", tc.score, got, tc.want)
		}
	}
}
