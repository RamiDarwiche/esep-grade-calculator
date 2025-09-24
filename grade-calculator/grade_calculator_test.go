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
	gradeCalculator.AddGrade("exam 1", 30, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 45, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 60, Assignment)
	gradeCalculator.AddGrade("exam 1", 62, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 66, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 73, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 78, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestNumericalGradeCalculation(t *testing.T) {
	expected_value := 71

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 72, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 74, Essay)

	actual_value := gradeCalculator.calculateNumericalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%d'; got '%d' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	expected_value := "assignment"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 40, Assignment)
	gradeCalculator.AddGrade("exam 1", 30, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 45, Essay)

	actual_value := gradeCalculator.grades[0].Type.String()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradePass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 80, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 93, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradePass2(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 85, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 93, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeFail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 70, Assignment)
	gradeCalculator.AddGrade("exam 1", 64, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 55, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeFail2(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 44, Assignment)
	gradeCalculator.AddGrade("exam 1", 33, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 21, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}