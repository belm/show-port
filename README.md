# 📡 Show-Port

A beautiful cross-platform command-line tool to display active network ports and their associated processes.

## ✨ Features

- 🌍 **Cross-platform**: Works on Linux, macOS, and Windows
- 🎨 **Beautiful UI**: Color-coded output with clear formatting
- 🔍 **Smart Display**: Aggregated view by default, eliminates duplicate connections
- 🏷️ **Service Recognition**: Automatically identifies common services (HTTP, MySQL, SSH, etc.)
- 📊 **Dual Modes**: Aggregated summary view or detailed connection view
- 🎯 **Flexible Filtering**: Filter by protocol, port, status, or connection count
- ⚡ **Fast & Efficient**: Lightweight and responsive
- 📦 **Easy Installation**: Simple `go install` command

## 🚀 Installation

### Using go install (recommended)

```bash
go install github.com/YOUR_USERNAME/show-port@latest
```

### Build from source

```bash
git clone https://github.com/YOUR_USERNAME/show-port.git
cd show-port
go build -o show-port
```

### Build for multiple platforms

```bash
# Build for all platforms (Linux, macOS, Windows)
make build-all

# Or manually for specific platform
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o show-port-linux-amd64 .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o show-port-windows.exe .
```

**Note**: Cross-compilation requires `CGO_ENABLED=0` because gopsutil uses CGO.

## 📖 Usage

### Basic Usage

Simply run the command:

```bash
show-port
```

**By default**, the tool shows:
- ✅ Only **LISTEN** ports (aggregated view)
- ✅ **Service names** for common ports (HTTP, HTTPS, MySQL, etc.)
- ✅ **Connection count** for each port
- ✅ Eliminates duplicate entries

This gives you a clean overview of what services are actually running on your system.

#### Aggregated View (Default)

```bash
show-port
```

Displays:
- **Protocol**: TCP, UDP, TCP6, UDP6 (color-coded)
- **Port**: Port number (highlighted in yellow)
- **Service**: Common service name (HTTP, MySQL, SSH, etc.)
- **Conns**: Number of connections on this port
- **PID**: Process ID
- **Process**: Process name (highlighted in cyan)
- **Status**: Port status (LISTEN, etc.)

#### Detailed View (All Connections)

```bash
show-port --all
```

Shows every connection with full details:
- **Protocol**, **Local Address**, **Port**, **Remote Address**
- **Status**, **PID**, **Process**, **Service**

### Command-Line Options

Show-port supports various filtering options:

```bash
# Show all connections (not aggregated)
show-port --all

# Show only listening ports (default behavior)
show-port --listen

# Show all TCP connections
show-port --all --protocol tcp --limit 20

# Find what's using port 3000
show-port --port 3000

# Show all ESTABLISHED connections
show-port --all --status ESTABLISHED

# Show top 10 listening ports
show-port --limit 10

# Combine filters - TCP listeners only
show-port --protocol tcp --limit 20

# Show version
show-port --version
```

#### Available Flags

- `--all`: Show all connections with full details (not aggregated)
- `--listen`: Show only ports in LISTEN state (default if no filter specified)
- `--protocol <type>`: Filter by protocol (tcp, udp, tcp6, udp6)
- `--status <status>`: Filter by status (LISTEN, ESTABLISHED, CLOSE_WAIT, etc.)
- `--port <number>`: Filter by specific port number
- `--limit <number>`: Limit the number of results (0 = no limit)
- `--version`: Display version information

#### 🎯 Common Use Cases

```bash
# Quick overview - what's listening on my machine?
show-port

# Is port 8080 in use?
show-port --port 8080

# What's using my database ports?
show-port --port 3306  # MySQL
show-port --port 5432  # PostgreSQL

# Show all active network connections
show-port --all --status ESTABLISHED

# Debug a specific service
show-port --port 3000 --all
```

### Example Output

#### Default View (Aggregated, Listen Only)

```
╔═══════════════════════════════════════════════════════════════╗
║           📡  Active Ports Monitor  📡                      ║
╚═══════════════════════════════════════════════════════════════╝

PROTOCOL   PORT     SERVICE              CONNS    PID      PROCESS         STATUS              
──────────────────────────────────────────────────────────────────────────────────────────────
tcp        22       SSH                  1        1234     sshd            LISTEN              
tcp        80       HTTP                 1        5678     nginx           LISTEN              
tcp        443      HTTPS                1        5678     nginx           LISTEN              
tcp        3000     Node.js/React Dev    1        9012     node            LISTEN              
tcp        3306     MySQL                1        3456     mysqld          LISTEN              
tcp        5432     PostgreSQL           1        7890     postgres        LISTEN              

✅ Total active connections: 6
```

#### Detailed View (--all)

```
PROTOCOL   LOCAL ADDR         PORT     REMOTE ADDR          STATUS       PID      PROCESS         SERVICE             
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
tcp        0.0.0.0           22       *                    LISTEN       1234     sshd            SSH                 
tcp        127.0.0.1         3000     *                    LISTEN       5678     node            Node.js/React Dev   
tcp        192.168.1.10      52341    142.250.185.46       ESTABLISHED  9012     chrome          -                   

✅ Total active connections: 3
```

## 🛠️ Requirements

- Go 1.20 or higher

## 📚 Dependencies

- [gopsutil](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system information
- [color](https://github.com/fatih/color) - Color output for terminal

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📝 License

MIT License - feel free to use this tool in your projects!

## 🐛 Troubleshooting

### Permission Issues

On some systems, you may need elevated privileges to view all process information:

- **Linux/macOS**: Run with `sudo show-port`
- **Windows**: Run as Administrator

### Port Information Not Showing

If you don't see expected ports, ensure:
1. The process is actually listening/connected
2. You have appropriate permissions
3. The port is bound to a network interface (not just internal IPC)

## ⭐ Show Your Support

Give a ⭐️ if this project helped you!

## 📧 Contact

Created by [YOUR_NAME] - feel free to reach out!

---

Made with ❤️ using Go
