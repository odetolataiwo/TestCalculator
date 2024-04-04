## Calculator Testing Framework Documentation
### Overview
The calculator testing framework is designed to facilitate the testing process, allowing developers to write and run both unit tests and integration tests efficiently. 
It provides a structured approach to organizing test cases, running tests, and reporting test results.

### Design
This testing framework is designed with modularity and extensibility in mind, allowing developers to easily add new test cases, test runners, and assertions as needed. 
It follows a modular architecture with the following components:

- Assertions: Contains functions for making assertions in tests, including comparing values, checking error conditions, etc.
- Runners: Provides test runners responsible for executing different types of tests, such as unit tests and integration tests. Each test runner encapsulates the logic for setting up test environments, running tests, and reporting test results.
- TestCases: Organizes test cases into separate directories for unit tests and integration tests. Each test case file contains test functions specific to the corresponding package or module.
- Testing: Serves as the entry point to the testing framework, providing initialization logic, configuration options, and utilities for managing tests.

### Structure
The testing framework directory structure is organized as follows:
```
testing/
│
├── assertions/
│   └── assertions.go
│
├── runners/
│   ├── unit_test_runner.go
│   ├── integration_test_runner.go
│   └── test_runner.go
│
├── testcases/
│   ├── unit/
│   │   └── calculator_test.go
│       └── parseoperand_test.go
│   └── integration/
│       └── calculator_integration_test.go
│
└── testing.go
```

## Usage
To use the testing framework, follow these steps:
To run unit tests, execute the following command:
```bash
go test -v ./...
```
or call this function in your code:
```go
testing.RunUnitTests()
```

To run integration tests, execute the following command:
```bash
go test -v -tags=integration ./...
```
or call this function in your code:
```go
testing.RunIntegrationTests()
```

## Adding New Tests
To add new test cases, follow these steps:
1. Create a new test case file in the appropriate directory under `testcases/`.
2. For unit tests, create a new file in the `unit/` directory. For integration tests, create a new file in the `integration/` directory.
3. Write test functions in the new test case file using the testing framework's assertions and utilities.
