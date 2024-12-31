# fynd

## Overview

`fynd` is a command-line utility inspired by `grep`. It allows users to search for patterns within files or standard input using regular expressions. Simple yet powerful, it provides a way to locate specific lines of text matching your regex patterns.

Additionally, I also included a **Lorem Ipsum Generator**, useful for generating random sentences for testing `fynd`. Feel free to tinker around with it.

---

## Features

- Search text files or standard input using regex.
- Outputs lines that match the given pattern.
- Compatible across macOS, Linux, and Windows.
- Includes a Lorem Ipsum Generator for testing.

---

## Installation

To install and make `fynd` work as a system-wide command:

### macOS / Linux

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```
2. **Build the binary**:
   ```bash
   go build -o fynd fynd.go
   ```
3. **Move the binary to a system-wide path**:
   ```bash
   sudo mv fynd /usr/local/bin/
   ```
4. **Verify installation**:
   ```bash
   fynd "pattern" sample.txt
   ```

### Windows

1. **Clone the repository**:
   ```powershell
   git clone <repository-url>
   cd <repository-folder>
   ```
2. **Build the binary**:
   ```powershell
   go build -o fynd.exe fynd.go
   ```
3. **Add the binary to PATH**:
   - Move `fynd.exe` to a directory included in your system's PATH, or update the PATH to include the directory where `fynd.exe` is located.
4. **Verify installation**:
   ```powershell
   fynd "pattern" sample.txt
   ```

---

## Usage

### `fynd`

```bash
fynd <pattern> [file]
```

- `<pattern>`: The regex pattern to search for.
- `[file]`: The file to search in. Use `-` to read from standard input.

**Example**:
```bash
fynd "error" log.txt
```

**Using Standard Input**:
```bash
echo "hello world" | fynd "hello"
```

---

### Lorem Ipsum Generator

The repository also includes a `lorem` program for generating random placeholder text.

#### Usage

```bash
go run lorem.go
```

**Features**:
- Generates random sentences with punctuation.
- Adds realistic pauses for text simulation.

**Example Output**:
```
Lorem ipsum dolor sit amet consectetur.
Sed do eiusmod tempor incididunt ut labore et dolore? Magna aliqua ut enim.
```

- Use it to test the `fynd` program

```
go run lorem.go | fynd "et"
```

---

## Contributing

Feel free to open issues or submit pull requests to enhance the functionality or fix any bugs. Contributions are always welcome!

---

## License

This project is licensed under the MIT License.

---

## Acknowledgements

`fynd` is inspired by the simplicity and power of Unix's `grep`. The Lorem Ipsum generator is included for developers and designers to simplify testing workflows.

