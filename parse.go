package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Events struct {
	User      string `json:"user"`
	Command   string `json:"command"`
	Service   string `json:"service"`
	Timestamp string `json:"timestamp"`
}

func main() {
	logs := []string{
		`{"user": "root", "command": "wget http://internal.repo", "service": "gce", "timestamp": "2025-04-20T11:30:00Z"}`,
		`{"user": "admin", "command": "wget http://safe.site", "service": "gce", "timestamp": "2025-04-20T11:10:00Z"}`,
		`{"user": "root", "command": "ls -la", "service": "gce", "timestamp": "2025-04-20T11:00:00Z"}`,
		`{"user": "root", "command": "wget http://stealth.download", "service": "gce", "timestamp": "2025-04-20T11:40:00Z"}`,
		`{"user": "bob", "command": "curl http://external", "service": "gce", "timestamp": "2025-04-20T11:50:00Z"}`,
	}

	fmt.Println(parseSusWget(logs))

}

func parseSusWget(log []string) []string {
	var susLogs []string

	cutOff := "2025-04-20T11:25:00Z"
	parsedCutoff, err := time.Parse(time.RFC3339, cutOff)
	if err != nil {
		fmt.Println("Error parsing time", err)
		return nil
	}

	for _, l := range log {
		var Temp Events
		err := json.Unmarshal([]byte(l), &Temp)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		timeStamp := Temp.Timestamp
		parsedTimeStamp, err := time.Parse(time.RFC3339, timeStamp)
		if err != nil {
			fmt.Println("Error parsing timestamp", err)
			continue
		}

		if Temp.User == "root" && strings.Contains(Temp.Command, "wget") && Temp.Service == "gce" && parsedTimeStamp.After(parsedCutoff) {
			msg := fmt.Sprintf("User = %v, Command = %v, Service = %v Timestamp = %v", Temp.User, Temp.Command, Temp.Service, Temp.Timestamp)
			susLogs = append(susLogs, msg)
		}

	}
	return susLogs
}
