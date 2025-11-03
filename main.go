package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

// PortInfo holds information about a port in use
type PortInfo struct {
	Protocol     string
	LocalAddr    string
	LocalPort    uint32
	RemoteAddr   string
	RemotePort   uint32
	Status       string
	PID          int32
	ProcessName  string
	ServiceName  string
	ConnCount    int // Number of connections for this port
}

const version = "1.0.0"

var (
	protocolFilter = flag.String("protocol", "", "Filter by protocol (tcp, udp, tcp6, udp6)")
	statusFilter   = flag.String("status", "", "Filter by status (LISTEN, ESTABLISHED, etc.)")
	portFilter     = flag.Int("port", 0, "Filter by specific port number")
	limitResults   = flag.Int("limit", 0, "Limit number of results (0 = no limit)")
	listenOnly     = flag.Bool("listen", false, "Show only LISTEN ports (default mode)")
	allConns       = flag.Bool("all", false, "Show all connections (not aggregated)")
	showVersion    = flag.Bool("version", false, "Show version information")
)

// Common service names by port number
var serviceMap = map[uint32]string{
	20:    "FTP-DATA",
	21:    "FTP",
	22:    "SSH",
	23:    "Telnet",
	25:    "SMTP",
	53:    "DNS",
	80:    "HTTP",
	110:   "POP3",
	143:   "IMAP",
	443:   "HTTPS",
	445:   "SMB",
	3000:  "Node.js/React Dev",
	3306:  "MySQL",
	5000:  "Flask/Dev Server",
	5432:  "PostgreSQL",
	6379:  "Redis",
	8000:  "Django/Dev",
	8080:  "HTTP-Alt/Proxy",
	8443:  "HTTPS-Alt",
	9000:  "PHP-FPM",
	27017: "MongoDB",
	3389:  "RDP",
	1433:  "MS-SQL",
	5900:  "VNC",
}

func main() {
	flag.Parse()

	// Show version if requested
	if *showVersion {
		fmt.Printf("show-port version %s\n", version)
		return
	}

	// Print header
	printHeader()

	// Get all network connections
	ports, err := getUsedPorts()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting port information: %v\n", err)
		os.Exit(1)
	}

	// If not showing all connections and no specific filters, default to listen-only
	if !*allConns && *statusFilter == "" && !*listenOnly {
		*listenOnly = true
	}

	// Apply filters
	ports = filterPorts(ports)

	// Aggregate connections by port if not showing all
	if !*allConns {
		ports = aggregatePorts(ports)
	}

	if len(ports) == 0 {
		fmt.Println("No active ports found matching the criteria.")
		return
	}

	// Display ports in a beautiful table
	displayPorts(ports)
	
	// Print summary
	printSummary(len(ports))
}

func printHeader() {
	cyan := color.New(color.FgCyan, color.Bold)
	fmt.Println()
	cyan.Println("╔═══════════════════════════════════════════════════════════════╗")
	cyan.Println("║           📡  Active Ports Monitor  📡                      ║")
	cyan.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Println()
}

func printSummary(count int) {
	green := color.New(color.FgGreen, color.Bold)
	fmt.Println()
	green.Printf("✅ Total active connections: %d\n\n", count)
}

func getUsedPorts() ([]PortInfo, error) {
	connections, err := net.Connections("all")
	if err != nil {
		return nil, err
	}

	var ports []PortInfo
	processCache := make(map[int32]string)

	for _, conn := range connections {
		// Only show connections with a local port
		if conn.Laddr.Port == 0 {
			continue
		}

		// Convert connection type to string
		protocol := getProtocolName(conn.Type)
		
		portInfo := PortInfo{
			Protocol:   protocol,
			LocalAddr:  conn.Laddr.IP,
			LocalPort:  conn.Laddr.Port,
			RemoteAddr: conn.Raddr.IP,
			RemotePort: conn.Raddr.Port,
			Status:     conn.Status,
			PID:        conn.Pid,
		}

		// Get process name
		if conn.Pid > 0 {
			if name, ok := processCache[conn.Pid]; ok {
				portInfo.ProcessName = name
			} else {
				if proc, err := process.NewProcess(conn.Pid); err == nil {
					if name, err := proc.Name(); err == nil {
						portInfo.ProcessName = name
						processCache[conn.Pid] = name
					}
				}
			}
		}

		if portInfo.ProcessName == "" {
			portInfo.ProcessName = "-"
		}

		// Add service name
		if serviceName, ok := serviceMap[portInfo.LocalPort]; ok {
			portInfo.ServiceName = serviceName
		} else {
			portInfo.ServiceName = "-"
		}

		portInfo.ConnCount = 1
		ports = append(ports, portInfo)
	}

	// Sort by port number
	sort.Slice(ports, func(i, j int) bool {
		if ports[i].LocalPort != ports[j].LocalPort {
			return ports[i].LocalPort < ports[j].LocalPort
		}
		return ports[i].Protocol < ports[j].Protocol
	})

	return ports, nil
}

func getProtocolName(connType uint32) string {
	switch connType {
	case 1:
		return "tcp"
	case 2:
		return "udp"
	case 3:
		return "tcp6"
	case 4:
		return "udp6"
	default:
		return "unknown"
	}
}

func aggregatePorts(ports []PortInfo) []PortInfo {
	// Group ports by local port number and protocol
	portMap := make(map[string]*PortInfo)
	
	for _, port := range ports {
		key := fmt.Sprintf("%s:%d", port.Protocol, port.LocalPort)
		
		if existing, ok := portMap[key]; ok {
			// Increment connection count
			existing.ConnCount++
			// Prefer LISTEN status if available
			if port.Status == "LISTEN" {
				existing.Status = "LISTEN"
				existing.PID = port.PID
				existing.ProcessName = port.ProcessName
			}
		} else {
			portCopy := port
			portMap[key] = &portCopy
		}
	}
	
	// Convert map back to slice
	var aggregated []PortInfo
	for _, port := range portMap {
		aggregated = append(aggregated, *port)
	}
	
	// Sort by port number
	sort.Slice(aggregated, func(i, j int) bool {
		if aggregated[i].LocalPort != aggregated[j].LocalPort {
			return aggregated[i].LocalPort < aggregated[j].LocalPort
		}
		return aggregated[i].Protocol < aggregated[j].Protocol
	})
	
	return aggregated
}

func filterPorts(ports []PortInfo) []PortInfo {
	var filtered []PortInfo
	
	for _, port := range ports {
		// Apply protocol filter
		if *protocolFilter != "" && !strings.EqualFold(port.Protocol, *protocolFilter) {
			continue
		}
		
		// Apply status filter
		if *statusFilter != "" && !strings.EqualFold(port.Status, *statusFilter) {
			continue
		}
		
		// Apply port filter
		if *portFilter > 0 && port.LocalPort != uint32(*portFilter) {
			continue
		}
		
		// Apply listen-only filter
		if *listenOnly && !strings.EqualFold(port.Status, "LISTEN") {
			continue
		}
		
		filtered = append(filtered, port)
		
		// Apply limit
		if *limitResults > 0 && len(filtered) >= *limitResults {
			break
		}
	}
	
	return filtered
}

func displayPorts(ports []PortInfo) {
	// Print table header manually with colors
	cyan := color.New(color.FgCyan, color.Bold)
	
	if *allConns {
		// Detailed view with all connections
		cyan.Printf("%-10s %-18s %-8s %-18s %-12s %-8s %-15s %-20s\n",
			"PROTOCOL", "LOCAL ADDR", "PORT", "REMOTE ADDR", "STATUS", "PID", "PROCESS", "SERVICE")
		fmt.Println("────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────")
	} else {
		// Aggregated view
		cyan.Printf("%-10s %-8s %-20s %-8s %-8s %-15s %-20s\n",
			"PROTOCOL", "PORT", "SERVICE", "CONNS", "PID", "PROCESS", "STATUS")
		fmt.Println("──────────────────────────────────────────────────────────────────────────────────────────────────────")
	}
	
	// Print rows with colors
	green := color.New(color.FgGreen)
	blue := color.New(color.FgHiBlue)
	yellow := color.New(color.FgYellow, color.Bold)
	magenta := color.New(color.FgMagenta)
	cyanNormal := color.New(color.FgCyan)
	
	for _, port := range ports {
		pidStr := "-"
		if port.PID > 0 {
			pidStr = strconv.Itoa(int(port.PID))
		}
		
		status := port.Status
		if status == "" {
			status = "-"
		}
		
		// Color code based on protocol
		protocolColor := color.New(color.FgWhite)
		switch port.Protocol {
		case "tcp", "tcp6":
			protocolColor = green
		case "udp", "udp6":
			protocolColor = blue
		}
		
		if *allConns {
			// Detailed view
			remoteAddr := port.RemoteAddr
			if remoteAddr == "" {
				remoteAddr = "*"
			}
			protocolColor.Printf("%-10s ", port.Protocol)
			fmt.Printf("%-18s ", port.LocalAddr)
			yellow.Printf("%-8s ", strconv.Itoa(int(port.LocalPort)))
			fmt.Printf("%-18s ", remoteAddr)
			fmt.Printf("%-12s ", status)
			magenta.Printf("%-8s ", pidStr)
			cyanNormal.Printf("%-15s ", port.ProcessName)
			fmt.Printf("%-20s\n", port.ServiceName)
		} else {
			// Aggregated view
			protocolColor.Printf("%-10s ", port.Protocol)
			yellow.Printf("%-8s ", strconv.Itoa(int(port.LocalPort)))
			fmt.Printf("%-20s ", port.ServiceName)
			magenta.Printf("%-8d ", port.ConnCount)
			fmt.Printf("%-8s ", pidStr)
			cyanNormal.Printf("%-15s ", port.ProcessName)
			fmt.Printf("%-20s\n", status)
		}
	}
}
