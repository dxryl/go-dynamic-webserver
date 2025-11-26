# Go vs Flask: Quick Comparison

**Bonus Section - Technology Comparison**  
*Moringa AI Capstone Project*

---

## Why Compare These Two?

I built my web server in **Go**, but I was curious how it compares to **Flask (Python)** - another popular choice for web servers. This comparison helped me understand the trade-offs between different technologies.

---

## Quick Overview

| Aspect | Go | Flask |
|--------|----|----|
| **Type** | Compiled programming language | Python web framework |
| **Created** | Google (2009) | Armin Ronacher (2010) |
| **Best For** | High-performance APIs, microservices | Rapid prototyping, data science apps |
| **Learning Curve** | Moderate | Easy |

---

## Side-by-Side Comparison

| Feature | Go | Flask | Winner |
|---------|----|----|--------|
| **Speed/Performance** | ⭐⭐⭐⭐⭐ (63K+ req/sec) | ⭐⭐⭐ (600-12K req/sec) | Go |
| **Development Speed** | ⭐⭐⭐ Moderate | ⭐⭐⭐⭐⭐ Very Fast | Flask |
| **Concurrency** | ⭐⭐⭐⭐⭐ Native (goroutines) | ⭐⭐ Limited (GIL) | Go |
| **Memory Usage** | ~5-10 MB | ~15-20 MB | Go |
| **Learning Curve** | Moderate (new syntax) | Easy (Python) | Flask |
| **Deployment** | Single binary file | Multiple files + dependencies | Go |
| **Ecosystem** | Growing | Massive (Python) | Flask |
| **Error Detection** | Compile-time | Runtime | Go |

---

## Code Comparison: Same Server, Different Languages

### Go Version (What I Built)

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

**Key Points:**
- 15 lines of code
- No external dependencies
- Just run: `go run main.go`

---

### Flask Version (Alternative)

```python
from flask import Flask
from datetime import datetime

app = Flask(__name__)

@app.route('/')
def handler():
    current_time = datetime.now().strftime("%I:%M %p %A %b %d %Y")
    return f"Hello! The current server time is: {current_time}"

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
```

**Key Points:**
- 12 lines of code
- Needs Flask: `pip install flask`
- Just run: `python main.py`

**Observation:** Flask is slightly shorter and more "Pythonic," but Go needs no dependencies.

---

## Performance Benchmarks

Based on industry benchmarks:

| Metric | Go | Flask (optimized) |
|--------|----|----|
| **Requests/Second** | 63,300 | 11,750 |
| **Average Latency** | <1ms | ~2ms |
| **Memory Usage** | 5 MB | 76 MB |
| **Concurrent Connections** | 10,000+ easily | Needs multi-processing |

**Bottom Line:** Go is roughly **5-6x faster** and uses **90% less memory**.

---

## When to Use Each

### Choose Go When:
✅ You need high performance (thousands of requests/sec)  
✅ Building microservices or APIs  
✅ Handling lots of concurrent connections  
✅ You want easy deployment (single binary)  
✅ Lower hosting costs matter  

**Real Companies Using Go:** Google, Uber, Docker, Netflix

---

### Choose Flask When:
✅ You need to build something fast (rapid prototyping)  
✅ Integrating with Python data science libraries  
✅ You're already comfortable with Python  
✅ Building small to medium traffic apps  
✅ You need the huge Python ecosystem  

**Real Companies Using Flask:** Netflix (internal tools), Pinterest (early versions)

---

## My Experience & Key Takeaways

### What I Learned About Go:
**Pros:**
- Incredibly fast and efficient
- Built-in web server (no extra setup)
- Catches errors before running (compile-time checking)
- One file to deploy (super easy)

**Cons:**
- Had to learn new syntax
- More verbose than Python
- Smaller community than Python

### If I Had to Choose Again:
- **For this project:** Go was perfect - I learned a valuable, in-demand skill
- **For quick prototypes:** I'd consider Flask for speed of development
- **For production systems:** Go all the way

---

## The Verdict

| Your Goal | Best Choice |
|-----------|-------------|
| Learn a high-performance language | Go |
| Build something in 1 day | Flask |
| Handle high traffic | Go |
| Integrate ML/data science | Flask |
| Minimize server costs | Go |
| Easiest to learn | Flask |

---

## Conclusion

Both are great, but they serve different purposes:

- **Go = Race car** 🏎️ (fast, efficient, built for performance)
- **Flask = Scooter** 🛴 (quick to start, easy to ride, gets you there)

For this capstone, Go taught me about performance-focused development and modern backend engineering. Flask would've been easier initially, but I wouldn't have learned as much about efficient server design.

**My recommendation:** Learn both! Start with Flask to understand web concepts quickly, then level up to Go when you need performance.

---

## Resources

**Go:**
- https://go.dev/tour/ - Interactive tutorial
- https://gobyexample.com/ - Learn by examples

**Flask:**
- https://flask.palletsprojects.com/ - Official docs
- https://blog.miguelgrinberg.com/post/the-flask-mega-tutorial-part-i-hello-world

**Benchmarks:**
- https://www.techempower.com/benchmarks/ - Real performance data

---

**Created by:** Noah Munda  
**Date:** November 26, 2025  
**Course:** Moringa AI Capstone Project - Bonus Section