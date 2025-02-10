# Golang 101

This repository contains a collection of fundamental concepts and examples in Go programming language (Golang). Each file demonstrates a specific concept, ranging from basic syntax to advanced topics such as concurrency, error handling, and testing.

## Folder Structure

```
golang-101
├─ 01_variables.go        # Variables declaration and usage
├─ 02_data_types.go       # Basic data types in Go
├─ 03_operators.go        # Operators in Go
├─ 04_control_flow.go     # Control flow statements (if, for, switch)
├─ 05_functions.go        # Functions in Go (declaration, parameters, return values)
├─ 06_closures.go         # Closures in Go
├─ 07_structs.go          # Structs and composition in Go
├─ 08_interfaces.go       # Interfaces in Go
├─ 09_pointers.go         # Pointers and memory management
├─ 10_goroutines.go       # Goroutines for concurrency
├─ 11_channels.go         # Channels for communication between goroutines
├─ 12_mutex.go            # Mutex for synchronization
├─ 13_error_handling.go   # Error handling in Go
├─ 14_panic_recover.go    # Panic and recover mechanisms
├─ 15_unit_test_test.go   # Unit tests in Go
├─ 16_benchmark_test.go   # Benchmark testing in Go
└─ README.md              # Project documentation
```

## Concept Overview

1. **Variables and Data Types**  
   Learn how to declare variables and understand the basic data types in Go such as `int`, `string`, `bool`, etc.

2. **Control Flow**  
   Demonstrates control flow statements such as `if`, `for`, and `switch` in Go.

3. **Functions**  
   Covers function declaration, passing parameters, returning values, and function scope.

4. **Closures**  
   Explains how closures work in Go, allowing functions to capture variables from their surrounding scope.

5. **Structs**  
   Introduces structs for organizing data and using composition to build complex types.

6. **Interfaces**  
   Shows how interfaces define behavior and how to implement them in Go.

7. **Pointers**  
   Discusses how pointers work in Go and how they are used to refer to memory addresses.

8. **Goroutines and Concurrency**  
   Explores how to use goroutines for concurrent programming in Go.

9. **Channels**  
   Explains how channels are used for communication between goroutines.

10. **Mutex for Synchronization**  
    Demonstrates how to use mutexes to prevent data races when multiple goroutines are accessing shared data.

11. **Error Handling**  
    Covers how to handle errors properly using Go's built-in error type.

12. **Panic and Recover**  
    Discusses how to handle panic situations and recover from errors during runtime.

13. **Unit Testing**  
    Introduces unit testing in Go, showing how to test functions with the `testing` package.

14. **Benchmark Testing**  
    Explains how to benchmark Go code using the `testing` package.

## How to Run the Examples

1. Clone the repository to your local machine:

    ```bash
    git clone https://github.com/your-username/golang-101.git
    ```

2. Navigate into the directory:

    ```bash
    cd golang-101
    ```

3. Run any `.go` file with the Go command:

    ```bash
    go run 01_variables.go
    ```

4. To run the tests and benchmarks:

    - Run unit tests:

      ```bash
      go test
      ```

    - Run benchmarks:

      ```bash
      go test -bench .
      ```

## Contributions

Feel free to fork this repository and contribute by submitting pull requests. If you find any issues or have suggestions, please open an issue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
