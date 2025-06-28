# 🔐 KeyWatch-Go

**KeyWatch-Go** is a real-time keystroke logger built with Go and WebSockets, designed strictly for **educational and ethical security research purposes**.

This project demonstrates how JavaScript can capture user keystrokes from a browser and send them to a Go backend using WebSockets. It also classifies whether the keystrokes are entered within form fields or elsewhere on the page—useful for simulating phishing pages or honeypots in controlled environments.

---

## 🚀 Features

- 📡 Real-time keystroke capture using WebSockets  
- 🧠 Classifies input as `form` or `other`  
- 🕒 Session-based logging with timestamps  
- 💻 Clean login page styled with HTML & CSS  
- 💾 Logs saved to `keystrokes.log` file

---

## ⚙️ How to Run Locally

> **Prerequisites:**  
> - [Go](https://golang.org/dl/) installed (v1.18+ recommended)

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/KeyWatch-Go.git
cd KeyWatch-Go
```

### 2. Run the Go server
```bash
go run main.go -listen-addr="127.0.0.1:8080" -ws-addr="127.0.0.1:8080"
```
### 3. Open the login form
```bash
http://127.0.0.1:8080
```
### 📁 File Structure
```bash
KeyWatch-Go/
├── main.go            # Go server logic
├── logger.js          # JavaScript for capturing keystrokes
├── static/
│   ├── form.html      # Main HTML form page
│   └── styles.css     # Page styling
├── keystrokes.log     # Generated log file
```
