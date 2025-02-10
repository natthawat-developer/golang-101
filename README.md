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
├─ 07_array_slice.go      # Arrays and Slices in Go
├─ 08_structs.go          # Structs and composition in Go
├─ 09_interfaces.go       # Interfaces in Go
├─ 10_pointers.go         # Pointers and memory management
├─ 11_maps.go             # Maps in Go
├─ 12_goroutines.go       # Goroutines for concurrency
├─ 13_channels.go         # Channels for communication between goroutines
├─ 14_select.go           # Select for multiple channels handling
├─ 15_mutex.go            # Mutex for synchronization
├─ 16_error_handling.go   # Error handling in Go
├─ 17_panic_recover.go    # Panic and recover mechanisms
├─ 18_unit_test_test.go   # Unit tests in Go
├─ 19_benchmark_test.go   # Benchmark testing in Go
└─ README.md              # Project documentation
```

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

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
