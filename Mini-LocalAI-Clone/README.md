# 🤖 Terminal-Based AI Chat Interface With GO

A powerful, intuitive terminal user interface for interacting with local AI models through Ollama. Features a beautiful TUI with chat history, multimodal support, and seamless model management.

---

## ✨ Features

- **🗨️ Beautiful Chat TUI** with persistent conversation history
- **💾 Local Storage** using SQLite - continue your conversations anytime
- **🎨 Intuitive Interface** powered by Bubble Tea
- **⚡ Multiple Model Support** - switch between different AI models seamlessly
- **🖼️ Multimodal Capabilities** - chat with AI about images (supported models only)
- **🎯 Visual Feedback** with spinners and formatted output
- **✏️ Customizable Prompts** for tailored responses

---

## 🚀 Getting Started

### Prerequisites

- Ollama installed on your system
- At least one model installed via Ollama (`ollama pull <model-name>`)
- For model options, visit the [Ollama Library](#)

### Installation

#### Build from Source

1. Clone the repository:
   ```bash
   git clone https://github.com/AkhilSharma90/GO-Native-LLM.git
   ```
   
2. Build the executable:
   ```bash
   go build
   ```
   
3. Run the application:
   ```bash
   ./FILENAME # On Unix-based systems
   ./FILENAME.exe # On Windows
   ```

---

## 🎮 Usage

### Interactive TUI Mode

Simply run the executable to launch the interactive TUI:
```bash
./executable
```

You'll be presented with options to:
- Start a new chat
- Continue previous conversations
- Manage your models
- Monitor running models

### Command Line Options

```bash
-v, --version    Print version information
-m, --manage     Manage installed models (update/delete)
-i, --install    Install new models
-r, --monitor    Monitor running model status
-h, --help       Show help message
```

### CLI Mode with Direct Input

```bash
# Using pipe
echo "Your text here" | ./executable --model="llama2" --prompt="Your prompt"

# Using input file
./executable --model="llama2" --prompt="Your prompt" < input.txt

# With image input (multimodal models only)
./executable --model="llava" --prompt="Describe this image" --images="path/to/image.png"
```

---

## ⌨️ Keyboard Shortcuts

### Chat Selection Screen

| Key       | Action                         |
|-----------|--------------------------------|
| ↑ / k      | Move up                        |
| ↓ / j      | Move down                      |
| → / l / pgdn | Next page                    |
| ← / h / pgup | Previous page               |
| g / home  | Go to start                     |
| G / end   | Go to end                       |
| enter     | Select chat                     |
| q         | Quit                            |
| d         | Delete chat                     |
| ctrl+n    | New chat                        |
| ?         | Toggle help                     |

### Main Chat Screen

| Key       | Action                         |
|-----------|--------------------------------|
| ctrl+up / k    | Scroll up                        |
| ctrl+down / j  | Scroll down                      |
| ctrl+u    | Half page up                     |
| ctrl+d    | Half page down                   |
| ctrl+p    | Previous message                 |
| ctrl+n    | Next message                     |
| ctrl+y    | Copy last response               |
| alt+y     | Copy selected message            |
| ctrl+o    | Toggle image picker              |
| ctrl+x    | Remove attachment                |
| ctrl+h    | Toggle help                      |
| ctrl+c    | Exit chat                        |

### Model Management Screens

| Key           | Action                         |
|---------------|--------------------------------|
| ↑ / k          | Move up                        |
| ↓ / j          | Move down                      |
| enter         | Select/Search/filter           |
| esc           | Clear filter                   |
| q / ctrl+c    | Quit                           |
| tab / n       | Next tab                       |
| shift+tab / p | Previous tab                   |
| u             | Update model                   |
| d             | Delete model                   |

---

## 🛠️ Technical Details

### Environment Configuration

- Default Ollama host: `http://localhost:11434`
- Can be configured using `OLLAMA_HOST` environment variable.

### Dependencies

Built with love using:
- [ollama](#) - Official Go SDK
- [bubbletea](#) - Terminal UI framework
- [glamour](#) - Markdown rendering
- [huh](#) - Terminal forms
- [lipgloss](#) - Terminal styling