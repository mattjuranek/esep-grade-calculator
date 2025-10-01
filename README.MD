# ESEP Grade Calculator

## About

This library is a simple utility to calculate a final grade from three categories:

- Assignment - 50%
- Exams - 35%
- Essays - 15%

It averages each category and then converts the weighted score to a letter grade (A–F).

## Getting Started

**Runtime requirements:**
Golang 1.24+

Clone your fork and enter the project directory:

```
git clone https://github.com/<username>/esep-grade-calculator.git
cd esep-grade-calculator
```

## How to Use

1. Create a GradeCalculator with **NewGradeCalculator()**
1. Add grades with **AddGrade(name, score, gradeType)** where gradeType is Assignment, Exam, or Essay
1. Call **GetFinalGrade()** to convert to the corresponding letter grade

## Testing

Run unit tests:

```
cd grade_calculator/
go test ./...
```

Run unit tests with code coverage:

```
cd grade_calculator/
go test -coverprofile=coverage.txt ./...
```
