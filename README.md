
# ♊ Zodiac

**Zodiac** is a Go-powered AI shell assistant that connects to Google Gemini and allows it to interact with your local machine. Gemini can write scripts, execute them, manage files, and run system commands — all through a secure interface you control.

> ✨ Powered by Go. Driven by Gemini.

---

## 🧠 Features

- 🔗 Connects to Google Gemini via API
- 📁 File system access (read/write/delete)
- ⚙️ System command execution
- 📝 Script generation and execution (e.g. Python, Bash, PowerShell)
- 🛡️ Sandbox-aware execution (WIP)
- 🚀 Cross-platform (Windows, macOS, Linux)
- 🧩 Plugin-friendly architecture (planned)

---

## 🛠️ Requirements

- **Go 1.21+**
- Google Gemini API key
- Python or shell installed (for script execution)

---

## 🚀 Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/zodiac.git
cd zodiac
```

### 2. Set Up Your API Key

Create a `.env` file or set the following environment variable:

```env
GEMINI_API_KEY=your_google_gemini_api_key
```

### 3. Build the App

```bash
go build -gcflags="all=-l -B" -ldflags="-s -w -extldflags '-static'" -trimpath
```

### 4. Run Zodiac

```bash
./Zodiac

# Or if using Windows

./Zodiac.exe
```

---

## 🧪 Example Usage

```txt
> gemini> Write a Python script to rename all .txt files in a folder to .bak
> Executing script...
> Done. Files renamed.
```

---

<!-- ## 📦 File Structure

```txt
zodiac/
├── cmd/             # CLI entry point
├── internal/
│   ├── gemini/      # Gemini API wrapper
│   ├── fs/          # File system helpers
│   ├── exec/        # Command/script execution logic
│   └── sandbox/     # (Optional) Script isolation
├── scripts/         # Temporary script cache
├── .env             # API key config
└── main.go
``` -->

---

## ⚠️ Security Notice

Zodiac gives an LLM (Gemini) **access to your local system**. Use caution:
- Limit execution to safe environments
- Avoid exposing Zodiac to the internet
- Review generated scripts before running (or implement a review mode)

---

<!-- ## 🗺️ Roadmap

- [ ] Command history and rollback
- [ ] Permission management (file/command white/blacklists)
- [ ] Web frontend (optional)
- [ ] Plugin system
- [ ] Logging and audit trail -->

---

## 📃 License

Licence pending. All rights reserved for now.

---

## 🌌 Credits

- [Gemini by Google](https://ai.google.dev)
- Built with ❤️ and Go
