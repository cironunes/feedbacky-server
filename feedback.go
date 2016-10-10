package main

type Feedback struct {
	Rating    int    `json:"rating"`
	Message   string `json:"message"`
	Questions string `json:"questions"`
}

type Feedbacks []Feedback
