# Fynd

## Overview

`fynd` is a command-line utility inspired by `grep`. It allows users to search for patterns within files or standard input using regular expressions. Simple yet powerful, it provides a way to locate specific lines of text matching your regex patterns.

Additionally, the repository includes a **Lorem Ipsum Generator**, useful for generating random sentences for testing text layouts or other placeholder needs.

---

## Features

- Search text files or standard input using regex.
- Outputs lines that match the given pattern.
- Compatible across macOS, Linux, and Windows.
- Includes a Lorem Ipsum Generator for testing.

---

## Installation

### macOS / Linux

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd <repository-folder>
   ```
2. **Build the binary**:
   ```bash
   go build -o /usr/local/bin/fynd fynd.go
   ```
3. **Verify installation**:
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

### Run Manually
You can pipe any log-producing command into `fynd`:

```bash
tail -f /var/log/system.log | fynd "error"
```

### Run System-Wide

#### **macOS** (Using `launchd`)
1. Create a `launchd` configuration file:
   ```bash
   sudo nano /Library/LaunchAgents/com.dhawal.fynd.plist
   ```

2. Add the following content:
   ```xml
   <?xml version="1.0" encoding="UTF-8"?>
   <plist version="1.0">
   <dict>
       <key>Label</key>
       <string>com.dhawal.fynd</string>
       <key>ProgramArguments</key>
       <array>
           <string>/usr/local/bin/fynd</string>
       </array>
       <key>RunAtLoad</key>
       <true/>
       <key>StandardInputPath</key>
       <string>/var/log/system.log</string>
       <key>StandardOutPath</key>
       <string>/var/log/fynd.log</string>
   </dict>
   </plist>
   ```

3. Load and start the service:
   ```bash
   sudo launchctl load -w /Library/LaunchAgents/com.dhawal.fynd.plist
   ```

4. Verify that `fynd` is running:
   ```bash
   launchctl list | grep com.dhawal.fynd
   ```

#### **Debian/Ubuntu** (Using `systemd`)
1. Create a systemd service file:
   ```bash
   sudo nano /etc/systemd/system/fynd.service
   ```

2. Add the following content:
   ```ini
   [Unit]
   Description=fynd Log Formatter
   After=network.target

   [Service]
   ExecStart=/usr/local/bin/fynd "pattern"
   Restart=always
   StandardInput=file:/var/log/syslog
   StandardOutput=file:/var/log/fynd.log

   [Install]
   WantedBy=multi-user.target
   ```

3. Enable and start the service:
   ```bash
   sudo systemctl enable fynd
   sudo systemctl start fynd
   ```

4. Verify that `fynd` is running:
   ```bash
   systemctl status fynd
   ```

5. Check logs:
   ```bash
   tail -f /var/log/fynd.log
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

**Use Case**:
- Useful for testing text layouts and placeholders in applications or designs.

---

## Contributing

Feel free to open issues or submit pull requests to enhance the functionality or fix any bugs. Contributions are always welcome!

---

## License

This project is licensed under the MIT License.

---

## Acknowledgements

`fynd` is inspired by the simplicity and power of Unix's `grep`. The Lorem Ipsum generator is included for developers and designers to simplify testing workflows.

