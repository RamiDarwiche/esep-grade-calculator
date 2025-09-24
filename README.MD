# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

# Grade Calculator

This project is built entirely in **Go** and relies only on the standard library.  
It provides functionality for calculating numerical and letter grades, as well as adding and averaging grades.

## Project Structure

- **README.md**  
  Briefly describes the codebase, its requirements, and the relevant commands a developer may want to run.

- **go.mod**  
  Defines the module for the project, including its identifier and Go version.  
  Used for dependency tracking if necessary.

- **grade_calculator.go**  
  The implementation of the grade calculator.  
  Contains all the logic for:

  - Grade types
  - Numerical grade calculations
  - Letter grade conversions
  - Grade addition and averaging

- **grade_calculator_test.go**  
  The testing suite for the grade calculator.  
  Includes test functions that validate:
  - Different final grade types
  - Expected letter grade results from various numerical inputs

Runtime requirements:
Golang 1.24+

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
