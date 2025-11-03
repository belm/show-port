# 📡 Show-Port

一个漂亮的跨平台命令行工具，用于显示活跃的网络端口及其关联进程。

## ✨ 特性

- 🌍 **跨平台**: 支持 Linux、macOS 和 Windows
- 🎨 **漂亮的界面**: 彩色输出，清晰的格式化
- 🔍 **智能显示**: 默认聚合视图，消除重复连接
- 🏷️ **服务识别**: 自动识别常见服务（HTTP、MySQL、SSH 等）
- 📊 **双模式**: 聚合摘要视图或详细连接视图
- 🎯 **灵活过滤**: 按协议、端口、状态或连接数过滤
- ⚡ **快速高效**: 轻量级且响应迅速
- 📦 **简易安装**: 简单的 `go install` 命令

## 🚀 安装

### 使用 go install（推荐）

```bash
go install github.com/YOUR_USERNAME/show-port@latest
```

### 从源码构建

```bash
git clone https://github.com/YOUR_USERNAME/show-port.git
cd show-port
go build -o show-port
```

### 构建多平台版本

```bash
# 构建所有平台（Linux、macOS、Windows）
make build-all

# 或手动构建特定平台
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o show-port-linux-amd64 .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o show-port-windows.exe .
```

**注意**: 交叉编译需要设置 `CGO_ENABLED=0`，因为 gopsutil 使用了 CGO。

## 📖 使用方法

### 基本用法

直接运行命令：

```bash
show-port
```

**默认情况下**，工具会显示：
- ✅ 仅显示 **监听中的端口**（聚合视图）
- ✅ 显示常见端口的 **服务名称**（HTTP、HTTPS、MySQL 等）
- ✅ 显示每个端口的 **连接数**
- ✅ 消除重复条目

这样可以清晰地了解系统上实际运行的服务。

#### 聚合视图（默认）

```bash
show-port
```

显示内容：
- **协议**: TCP、UDP、TCP6、UDP6（带颜色编码）
- **端口**: 端口号（黄色高亮）
- **服务**: 常见服务名称（HTTP、MySQL、SSH 等）
- **连接数**: 该端口上的连接数量
- **PID**: 进程 ID
- **进程**: 进程名称（青色高亮）
- **状态**: 端口状态（LISTEN 等）

#### 详细视图（所有连接）

```bash
show-port --all
```

显示每个连接的完整详情：
- **协议**、**本地地址**、**端口**、**远程地址**
- **状态**、**PID**、**进程**、**服务**

### 命令行选项

Show-port 支持多种过滤选项：

```bash
# 显示所有连接（不聚合）
show-port --all

# 仅显示监听端口（默认行为）
show-port --listen

# 显示所有 TCP 连接
show-port --all --protocol tcp --limit 20

# 查找使用端口 3000 的程序
show-port --port 3000

# 显示所有已建立的连接
show-port --all --status ESTABLISHED

# 显示前 10 个监听端口
show-port --limit 10

# 组合过滤 - 仅 TCP 监听端口
show-port --protocol tcp --limit 20

# 显示版本信息
show-port --version
```

#### 可用参数

- `--all`: 显示所有连接的完整详情（不聚合）
- `--listen`: 仅显示 LISTEN 状态的端口（未指定过滤器时的默认行为）
- `--protocol <类型>`: 按协议过滤（tcp、udp、tcp6、udp6）
- `--status <状态>`: 按状态过滤（LISTEN、ESTABLISHED、CLOSE_WAIT 等）
- `--port <数字>`: 按特定端口号过滤
- `--limit <数字>`: 限制结果数量（0 = 无限制）
- `--version`: 显示版本信息

#### 🎯 常见使用场景

```bash
# 快速查看 - 我的机器上在监听什么？
show-port

# 端口 8080 被占用了吗？
show-port --port 8080

# 查看数据库端口使用情况
show-port --port 3306  # MySQL
show-port --port 5432  # PostgreSQL

# 显示所有活跃的网络连接
show-port --all --status ESTABLISHED

# 调试特定服务
show-port --port 3000 --all
```

### 输出示例

#### 默认视图（聚合，仅监听）

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

#### 详细视图（--all）

```
PROTOCOL   LOCAL ADDR         PORT     REMOTE ADDR          STATUS       PID      PROCESS         SERVICE             
────────────────────────────────────────────────────────────────────────────────────────────────────────────────────
tcp        0.0.0.0           22       *                    LISTEN       1234     sshd            SSH                 
tcp        127.0.0.1         3000     *                    LISTEN       5678     node            Node.js/React Dev   
tcp        192.168.1.10      52341    142.250.185.46       ESTABLISHED  9012     chrome          -                   

✅ Total active connections: 3
```

## 🛠️ 系统要求

- Go 1.20 或更高版本

## 📚 依赖库

- [gopsutil](https://github.com/shirou/gopsutil) - 跨平台的进程和系统信息获取库
- [color](https://github.com/fatih/color) - 终端彩色输出库

## 🤝 贡献

欢迎贡献！请随时提交 Pull Request。

## 📝 许可证

MIT License - 可自由在您的项目中使用！

## 🐛 故障排除

### 权限问题

在某些系统上，您可能需要提升权限才能查看所有进程信息：

- **Linux/macOS**: 使用 `sudo show-port` 运行
- **Windows**: 以管理员身份运行

### 进程名称显示为 "-"

这通常是权限问题。某些进程的信息需要管理员/root 权限才能访问。使用 `sudo` 或以管理员身份运行可以解决。

### 端口信息未显示

如果看不到预期的端口，请确保：
1. 进程确实在监听/连接
2. 您有适当的权限
3. 端口绑定到网络接口（而不仅是内部 IPC）

## 💡 提示

- **默认视图**最适合快速查看系统上运行的服务
- 使用 `--all` 查看详细的网络连接信息
- 使用 `--limit` 限制输出，便于快速浏览
- 组合多个过滤器以精确定位您需要的信息

## ⭐ 支持项目

如果这个项目对您有帮助，请给个 ⭐️！

---

使用 Go 和 ❤️ 制作
