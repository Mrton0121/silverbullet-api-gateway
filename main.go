package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	sbapi "github.com/Mrton0121/silverbullet-go-api"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Healthcheck", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Read the "data" field
	data := r.FormValue("data")
	fmt.Println("[LOG]: Received data: %s\n", data) // Logs the received data

	// Getting the SilverBullet variables
	endpoint := os.Getenv("SB_URL")
	token := os.Getenv("SB_TOKEN")
	page := os.Getenv("SB_PAGE")

	fmt.Printf("[DEBUG]: SB_URL=%s, SB_TOKEN=%s, SB_PAGE=%s\n", endpoint, token, page)

	// Getting data pattern
	data_pattern := os.Getenv("DATA_PATTERN")

	// Get separator value, it's a newline on default
	separator := os.Getenv("SEPARATOR")
	if separator == "" {
		separator = "\n"
	}

	// Create sbapi client
	sbClient := sbapi.NewClient(endpoint, token)

	// Replace magic variables
	if data_pattern != "" {
		data = strings.Replace(data_pattern, "[TEXT]", data, -1)
		data = strings.Replace(data, "[DATE]", time.Now().Format(time.DateTime), -1)
		data = strings.Replace(data, "[SEPARATOR]", separator, -1)
		data = strings.Replace(data, "[TAB]", "\t", -1)
	}

	fmt.Printf("[LOG]: Appending %s with: %s%s\n", page, data, separator)

	resp, err := sbClient.Get(page)
	if err != nil {
		fmt.Println(err.Error())
	}

	if resp != "404 Not Found" {
		resp, err = sbClient.Append(page, data, separator)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else {
		resp, err = sbClient.Put(page, data)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("[LOG]: New page created: %s\n", resp)

		resp, err = sbClient.Append(page, separator, "")
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("[LOG]: Added separator to new page: %s\n", resp)
	}

	fmt.Printf("[LOG]: Appending %s: %s", page, resp)

	// Send a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Data received\n"))
}

func main() {

	// Getting the SilverBullet variables
	endpoint := os.Getenv("SB_URL")
	token := os.Getenv("SB_TOKEN")
	page := os.Getenv("SB_PAGE")

	fmt.Printf("[DEBUG]: SB_URL=%s, SB_TOKEN=%s, SB_PAGE=%s\n", endpoint, token, page)

	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
