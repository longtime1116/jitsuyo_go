package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Comment struct {
	Message  string
	UserName string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	var mutex = &sync.RWMutex{}

	// curl http://localhost:8080/
	http.HandleFunc("/", helloHandler)

	// curl http://localhost:8080/comments
	// curl http://localhost:8080/comments
	// curl -X POST http://localhost:8080/comments -H "Content-Type: application/json" -d '{"Message": "test1", "UserName": "test_user1"}'

	comments := make([]Comment, 0, 100)
	comments = append(comments, Comment{Message: "1st message", UserName: "Sample "})
	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			// 本来はDBから読み取るが、ここではそうでないので、読み取り時の書き込みをケアしてロックしておく
			mutex.RLock()

			if err := json.NewEncoder(w).Encode(comments); err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s}`, err), http.StatusInternalServerError)
				return
			}
			mutex.RUnlock()
		case http.MethodPost:
			var c Comment
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status": "%s}`, err), http.StatusInternalServerError)
				return
			}
			mutex.Lock()
			comments = append(comments, c)
			mutex.Unlock()
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status": "created"}`))
		default:
			http.Error(w, `{"status": "permits only GET or POST"}`, http.StatusInternalServerError)
		}

	})

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
