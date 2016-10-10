package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func FeedbackIndex(w http.ResponseWriter, r *http.Request) {
	feedbacks := Feedbacks{
		Feedback{Rating: 4, Message: "Great talk", Questions: "Do you recommend TS?"},
		Feedback{Rating: 3, Message: "Nice talk", Questions: "Do you recommend Angular 2?"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(feedbacks); err != nil {
		panic(err)
	}
}
