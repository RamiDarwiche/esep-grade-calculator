package esepunittests

type GradeCalculator struct {
	grades []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		grades: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetPassFail() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 70 {
		return "Pass"
	}

	return "Fail"
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
		gc.grades = append(gc.grades, Grade{
			Name:  name,
			Grade: grade,
			Type:  gradeType,
		})
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignmentAverage := computeAverage(filterGrades(gc.grades, Assignment))
	examAverage := computeAverage(filterGrades(gc.grades, Exam))
	essayAverage := computeAverage(filterGrades(gc.grades, Essay))

	weightedGrade := float64(assignmentAverage)*0.5 +
		float64(examAverage)*0.35 +
		float64(essayAverage)*0.15

	return int(weightedGrade)
}

func filterGrades(grades []Grade, gradeType GradeType) []Grade {
	filtered := make([]Grade, 0)
	for _, g := range grades {
		if g.Type == gradeType {
			filtered = append(filtered, g)
		}
	}
	return filtered
}
func computeAverage(grades []Grade) int {
	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}

	return sum / len(grades)
}
