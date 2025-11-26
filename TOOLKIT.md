# Go Dynamic Web Server - Beginner's Toolkit

## 1. Title & Objective

**Technology Chosen:** GoLang (Go)

**Why I Chose This:** I wanted to learn a modern, fast programming language that's highly valued in the tech industry. Go is known for its simplicity and efficiency in building web servers, and since I'm already proficient in Java, learning Go would expand my backend development skills. Plus, Go's growing popularity at companies like Google and Uber made it an attractive choice.

**End Goal/Objective:** This toolkit teaches beginners how to use Go (Golang) to build a simple but functional web server that serves dynamic content (current server time). It demonstrates how to learn new technology using Generative AI prompts, including installation, setup, coding, debugging, and documenting the process.

---

## 2. Quick Summary of the Technology

**What is Go?**

Go (Golang) is a modern, open-source programming language created at Google. It emphasizes:
- Simplicity
- High performance
- Easy concurrency
- Reliability

It's ideal for backend development, cloud services, and network applications.

**Where is it used?**

Go is used by companies like Google, Uber, Docker, and Netflix for building high-performance backend services.

**Real-world Example:**

Docker (the popular containerization platform) is written entirely in Go because of its speed and efficiency in handling concurrent operations. This shows that Go is capable of powering huge and complex systems.

---

## 3. System Requirements

- **Operating System:** Windows 10/11
- **Go Version:** 1.21 or higher
- **Tools Needed:**
  - Command Prompt or PowerShell
  - Text Editor (Notepad, VS Code, or any code editor)
  - Web Browser (Chrome, Firefox, Edge, etc.)

---

## 4. Installation & Setup Instructions

### Step 1: Install Go

1. Go to https://go.dev/dl/
2. Download the Windows installer (.msi file)
3. Run the installer and follow the setup wizard
4. Accept default installation path (`C:\Program Files\Go`)
5. Restart Your PC - This helps Windows register Go's PATH

### Step 2: Verify Installation

Open Command Prompt and run:
```bash
go version
```

You should see output like: `go version go1.21.x windows/amd64`

If this appears, Go is installed successfully.

### Step 3: Create the Project Folder

1. Create a folder at: `C:\Users\<yourname>\go-webserver`

2. Open Command Prompt and navigate to that folder:
```bash
cd C:\Users\<yourname>\go-webserver
```

3. Initialize the Go module:
```bash
go mod init go-webserver
```

### Step 4: Create main.go

1. Open File Explorer and navigate to your `go-webserver` folder
2. Right-click in the empty space → **New** → **Text Document**
3. You'll see `New Text Document.txt` appear
4. Click on it once to select it, then press **F2** (or right-click → Rename)
5. **IMPORTANT:** Delete everything and type: `main.go`
   - Make sure there's NO `.txt` at the end
   - Windows might warn you about changing file extensions - click **Yes**
6. Double-click `main.go` to open it with Notepad
7. Copy and paste the code from section 5 below
8. Save the file (**Ctrl + S**)

---

## 5. Minimal Working Example

Copy this code and paste it into `main.go`:

### Code: main.go
```go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("3:04 PM Monday Jan 2 2006")
    message := fmt.Sprintf("Hello! The current server time is: %s", currentTime)
    fmt.Fprintf(w, message)
}

func main() {
    http.HandleFunc("/", handler)

    fmt.Println("Starting server on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

```

### Expected Output

When you run:
```bash
go run main.go
```

You should see:
```
Starting server on http://localhost:8080
```
Visit: http://localhost:8080
**In the browser (http://localhost:8080):**

It will display dynamic content such as:
```
Hello! The current server time is: 4:58 PM Tuesday Jan 23 2025
```

### Brief Explanation

This code creates a simple HTTP web server using Go's built-in `net/http` package. Here's what each part does:

1. **`handler` function**: Responds to HTTP requests by getting the current server time using `time.Now()`, formatting it nicely, and sending it back to the browser.

2. **`main` function**: Sets up the server to listen on port 8080 and routes all requests ("`/`") to the `handler` function.

3. **Dynamic content**: Every time you refresh the browser, `time.Now()` is called again, showing the updated server time.

The result is a fully functional web server with dynamic content using only Go's standard library—no external dependencies needed!

---

## 6. AI Prompt Journal

### Prompt 1: Understanding Go

**Prompt:** "I'm proficient in Java and want to learn Golang. Could you help me create a structured learning journey plan with: - 3-4 distinct learning phases - Prerequisites for each phase - 4-5 specific learning steps per phase - Verification activities for each phase"

**Response Summary:** The AI provided a structured Go learning plan in four phases. Phase 1 covers Go basics like syntax, types, error handling, and structs/methods. Phase 2 focuses on interfaces, concurrency, pointers, and Go's tools, with verification through small programs and tests. Phase 3 is about building real-world applications using HTTP servers, frameworks, databases, and idiomatic project structure. Phase 4 covers production-level skills, including advanced concurrency, profiling, testing, architecture, and deployment, verified by optimizing, testing, and containerizing applications.

**My Evaluation:** The content matched my request for a structured journey, highlighting both language fundamentals and practical applications.

---

### Prompt 2: Installation Help

**Prompt:** "How do I install Go on Windows?"

**Response Summary:** The AI provided a step-by-step guide to installing Go on Windows. It included downloading the official MSI installer, running it with default settings, checking that the PATH is set correctly, and verifying the installation with go version and go env. It also suggested setting up a workspace folder and testing with a simple "Hello, Go!" program to ensure everything works.

**My Evaluation:** The instructions were easy to follow, fully relevant, and practical for immediate use. They included useful verification steps and optional workspace setup, making the process beginner-friendly.

---

### Prompt 3: Building the Server

**Prompt:** "I'd like to do a Mini Web Server - Serve dynamic content (current server time) as my Go project example."

**Response Summary:** AI gave me a minimal Go HTTP server using net/http, and explained each function.

**My Evaluation:** This guided the foundation of my mini project.

---

### Prompt 4: Adding Dynamic Content

**Prompt:** "How do I add dynamic content (like the current time) to a Go web server?"

**Response Summary:** AI explained how to import time and format timestamps, and how to insert them into HTTP responses.

**My Evaluation:** This was essential for the project goal — serving dynamic content.

---

## 7. Common Issues & Fixes

### Issue 1: "go: command not found"

**Problem:** Go is not installed or not in PATH

**Solution:** 
- Reinstall Go
- Restart Command Prompt
- Verify with `go version`

---

### Issue 2: "cannot find module providing package"

**Problem:** Missing go.mod file causes this

**Solution:** Run `go mod init go-webserver` in your project folder

---

### Issue 3: Port 8080 Already in Use

**Problem:** Error message: "address already in use" or "bind: Only one usage of each socket address"

**Solution:** 
- Another program is using port 8080
- Either stop that program, or change the port in your code to `:8081` or `:8090`
- On Windows, you can find what's using port 8080 with: `netstat -ano | findstr :8080`

---

## 8. References

**Official Documentation:**
- Go Official Website: https://go.dev/
- Go Documentation: https://go.dev/doc/
- Go Tour (Interactive Tutorial): https://go.dev/tour/

**Helpful Tutorials:**
- https://go.dev/doc/articles/wiki/
- https://gobyexample.com/
- https://go.dev/doc/install

**AI Tools Used:**
- Claude AI (ai.moringaschool.com)
- ChatGPT

---

## Conclusion

Through this project, I learned how Go's simplicity and built-in packages make it powerful for web development. Using AI prompts (Claude and ChatGPT) significantly accelerated my learning by providing structured guidance, instant explanations, and debugging help. The AI helped me understand Go's syntax, install it correctly, and build a working web server in just a few days—something that would have taken much longer with traditional learning methods alone. This toolkit demonstrates that combining AI-powered learning with hands-on coding is an effective approach to mastering new technologies quickly.

---

**Created by:** Noah Munda  
**Date:** November 26, 2025  
**Course:** Moringa AI Capstone Project

