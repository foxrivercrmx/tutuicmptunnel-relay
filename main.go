package main

import (
	"encoding/json"
	"fmt"
	"flag"
	"log"
	"net"
	"net/http"
	"os/exec"
	"strings"
)

type RegisterPayload struct {
	PSK        string `json:"psk"`
	UID        string `json:"uid"`
	ClientIP   string `json:"client_ip"`
	ClientPort string `json:"client_port"`
	Comment    string `json:"comment"`
}

func formatListenAddress(input string) string {
	input = strings.TrimSpace(input)

	if input == "" {
		return "0.0.0.0:8080"
	}
	if !strings.Contains(input, ":") {
		return "0.0.0.0:" + input
	}

	host, port, err := net.SplitHostPort(input)
	if err != nil {
		return ""
	}
	if host == "" {
		host = "0.0.0.0"
	}

	return net.JoinHostPort(host, port)
}

func main() {
	apiPath := flag.String("path", "/api/register-tuctl", "Path for API request")
	presharedkey := flag.String("psk", "your-secret-psk", "PSK must same with client or GUI. Strong PSK is recommended")
	modeCtl := flag.String("mode", "tuctl", "tutuicmptunnel mode. tuctl or ktuctl or absolute path")
	inputAddr := flag.String("listen", "127.0.0.1:8080", "Listening address. IP:PORT")
	flag.Parse()

	http.HandleFunc(*apiPath, func(w http.ResponseWriter, r *http.Request) {
		var payload RegisterPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if payload.PSK != *presharedkey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		args := []string{"server-add", "uid", payload.UID, "address", payload.ClientIP, "port", payload.ClientPort, "comment", payload.Comment}
		
		cmd := exec.Command(*modeCtl, args...)
		output, err := cmd.CombinedOutput()
		
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to register: %v\nOutput: %s", err, output), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Peer successfully registered!"))
	})

	listenAddress := formatListenAddress(*inputAddr)
	if listenAddress == "" {
		log.Fatalf("Error: Format listen address '%s' not valid!", *inputAddr)
	}

	fmt.Printf("[+] Backend running...\n")
	fmt.Printf("[*] Listening on : %s\n", listenAddress)
	fmt.Printf("[*] API Path     : %s\n", *apiPath)
	fmt.Printf("[*] Expected PSK : %s\n", *presharedkey)
	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		log.Fatalf("Failed: %v", err)
	}
}