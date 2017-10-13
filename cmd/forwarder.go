package main

import (
	"io"
    "log"
	"net/http"
	"os"
	"strconv"
)

var httpClient = &http.Client{
}

func forward(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Query().Get("path")
	if len(path) == 0 {
		io.WriteString(w, "Provide 'path' query param to which request should be forwarded")
		return 
	}

	newReq, err := http.NewRequest(r.Method, path, nil)
	if err != nil {
		io.WriteString(w, "Could not create request to " + path + ". Reason: " + err.Error())
		return
	}
	log.Println("request created to: " + path)
	
	resp, err := httpClient.Do(newReq)
	if err != nil {
		io.WriteString(w, "Could not send request to " + path + ". Reason: " + err.Error())
		return
	}
	defer resp.Body.Close()

	io.Copy(w, resp.Body)
}


func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}

	_, err := strconv.Atoi(port);
	if err != nil {
		panic("PORT variable must be number or empty( will default to 8080) ")
	}

	http.HandleFunc("/", forward)

	log.Println("Will start listen on " + port)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}