🔐 KeyWatch-Go
KeyWatch-Go is a real-time keystroke logger built with Go and WebSockets, designed strictly for educational and ethical security research purposes.

This project demonstrates how JavaScript can capture user keystrokes from a browser and send them to a Go backend using WebSockets. It also classifies whether the keystrokes are entered within form fields or elsewhere on the page—useful for simulating phishing pages or honeypots in controlled environments.

🚀 Features
📡 Real-time keystroke capture using WebSockets

🧠 Classifies input as form or other

🕒 Session-based logging with timestamps

💻 Clean login page styled with HTML & CSS

💾 Logs saved to keystrokes.log file

⚙️ How to Run Locally
Prerequisites:

Go installed (v1.18+ recommended)

1. Clone the repository
bash
Copy
Edit
git clone https://github.com/yourusername/KeyWatch-Go.git
cd KeyWatch-Go
2. Run the Go server
bash
Copy
Edit
go run main.go -listen-addr="127.0.0.1:8080" -ws-addr="127.0.0.1:8080"
This starts:

The HTML frontend at http://127.0.0.1:8080

The JavaScript logger at http://127.0.0.1:8080/k.js

WebSocket endpoint at /ws

3. Open the login form
Navigate to:
http://127.0.0.1:8080
Start typing into the form or anywhere on the page — keystrokes will be logged to keystrokes.log.

📁 File Structure
graphql
Copy
Edit
KeyWatch-Go/
├── main.go            # Go server logic
├── logger.js          # JavaScript for capturing keystrokes
├── static/
│   ├── form.html      # Main HTML form page
│   └── styles.css     # Page styling
├── keystrokes.log     # Generated log file
⚠️ Disclaimer
This tool is for educational purposes only.
Do not use it on websites or systems you do not own or have explicit permission to test. Unauthorized use may violate laws or platform policies.
