package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "bytes"
    "encoding/json"
)


func TestHandlerMethodNotAllowed(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusMethodNotAllowed{
        t.Errorf("Wrong status code: got %v want %v",
            status, http.StatusMethodNotAllowed)
    }
}


func sendJsonToServer(t *testing.T, dataServer *DataServer) {
    jsonValue, _ := json.Marshal(dataServer)
    req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handler)
    handler.ServeHTTP(rr, req)

    // Check the status code is what we expect.
    if status := rr.Code; status != http.StatusOK{
        t.Errorf("Wrong status code: Got %v, expected %v",
            status, http.StatusOK)
    }
    if _, exist := sessions[dataServer.SessionId]; !exist {
        t.Errorf("Data was not created in the hash.")
    }

}


func TestCopyAndPasteData(t *testing.T) {
    dataServer := DataServer{EventType: "copyAndPaste", SessionId: "12", WebsiteUrl: "http://testweb.com", InputId: "nameInput"}
    sendJsonToServer(t, &dataServer)
}


func TestResizedData(t *testing.T) {
    resizeFrom := Dimension{Width: "10px", Height: "10px"}
    dataServer := DataServer{EventType: "screenResized", SessionId: "111", WebsiteUrl: "http://testweb.com", ResizeFrom: resizeFrom, ResizeTo: resizeFrom}
    sendJsonToServer(t, &dataServer)
}

func TestTimeTakendData(t *testing.T) {
    dataServer := DataServer{EventType: "timeTaken", SessionId: "121", WebsiteUrl: "http://testweb.com", FormCompletionTime: 123}
    sendJsonToServer(t, &dataServer)
}