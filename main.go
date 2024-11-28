package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type TimeResponse struct {
	CurrentTime string `json:"current_time"`
}

var db *sql.DB

func init() {
	// MySQL credentials
	dsn := "root:#Mysql@123@tcp(localhost:3306)/go_api"
	var err error
	// Open a connection to the database
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging the database: ", err)
	}
}

func main() {
	http.HandleFunc("/current-time", func(w http.ResponseWriter, r *http.Request) {
		// Load the Toronto timezone
		loc, err := time.LoadLocation("America/Toronto")
		if err != nil {
			http.Error(w, "Failed to load timezone", http.StatusInternalServerError)
			log.Printf("Error loading timezone: %v", err)
			return
		}

		// Get the current time in Toronto and format it
		currentTime := time.Now().In(loc).Format("2006-01-02 15:04:05")

		// Insert the current time into the time_log table
		_, err = db.Exec("INSERT INTO time_log (timestamp) VALUES (?)", currentTime)
		if err != nil {
			http.Error(w, "Failed to insert time into database", http.StatusInternalServerError)
			log.Printf("Error inserting time into database: %v", err)
			return
		}

		// Create a response with the formatted time
		response := TimeResponse{CurrentTime: currentTime}

		// Set the content type to JSON and send the response
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
			log.Printf("Error encoding JSON: %v", err)
		}
	})

	// Start the web server
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
