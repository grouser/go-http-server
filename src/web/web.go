package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)

// Store all the sessions created
var sessions = make(map[string]Data)

type DataServer struct{
	// Match the json messages from the server
	EventType string `json:"eventType"`
	SessionId  string `json:"sessionId"` 
	ResizeFrom         Dimension `json: resizeFrom`
	ResizeTo           Dimension `json: resizeTo`
	Pasted bool `"json:pasted"`
	InputId string `json:"inputId"` // map[fieldId]true
	FormCompletionTime int `json:"time"`// Seconds
	WebsiteUrl string `json: websiteUrl`  
}

type Data struct {
	SessionId  string 
	ResizeFrom         Dimension 
	ResizeTo           Dimension 
	CopyAndPaste       map[string]bool  // map[fieldId]true
	FormCompletionTime int // Seconds
	WebsiteUrl string 
}

type Dimension struct {
	Width  string  `json: "width"`
	Height string  `json: "height"`  
}

func handler(w http.ResponseWriter, r *http.Request){
	// These headers are for testing purposes. It will make possible to use the index.html without a web server
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	switch r.Method {
		case "OPTIONS":
			// Neccesary for enabling CORS 
			w.WriteHeader(http.StatusOK)
			return
		case http.MethodPost:
			var data DataServer
			if r.Body == nil {
				http.Error(w, "Please send a request body", 400)
				return
			}
			err := json.NewDecoder(r.Body).Decode(&data) 
			if err != nil {
				http.Error(w, err.Error(), 400)
				return
			}
			createDataFromDataServer(&data);
			fmt.Fprintf(w, "")
		default:
			fmt.Printf(http.MethodPost)
			http.Error(w, http.MethodPost, 405)
			return
	}
	
}

func createDataFromDataServer(dataServer *DataServer) {
	// Insert the data from the server into the sessions hash.
	data, isExist := sessions[dataServer.SessionId]
	if !isExist {
		// If map is nill, it will be initialised
		data.CopyAndPaste = make(map[string]bool)
	}
	data.SessionId = dataServer.SessionId
	data.WebsiteUrl = dataServer.WebsiteUrl
	
	switch dataServer.EventType {
		case "copyAndPaste":
			data.CopyAndPaste[dataServer.InputId] = dataServer.Pasted
		case "timeTaken":
			// We could delete the key from the sessions hash, as the form has been submitted
			data.FormCompletionTime = dataServer.FormCompletionTime
		case "screenResized":
			data.ResizeFrom = dataServer.ResizeFrom
			data.ResizeTo = dataServer.ResizeTo
	}
	
	// Print the current state of the struct
	fmt.Printf("%+v\n", data)
	// Save data into the hash. 
	sessions[data.SessionId] = data
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
