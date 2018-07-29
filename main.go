package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {

	quizService.Init()

	router := mux.NewRouter()

	router.HandleFunc("/questions", getQuestions).Methods("GET")
	router.HandleFunc("/answers", insertAnswers).Methods("POST", "OPTIONS")
	router.HandleFunc("/users/{id}/results", getQuizResults).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, HEAD, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding")
}

func getQuestions(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	questions, err := quizService.GetQuestions()

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(questions)
}

func insertAnswers(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	if (*r).Method == "OPTIONS" {
		return
	}

	var userAnswerContainer *UserAnswerContainer

	_ = json.NewDecoder(r.Body).Decode(&userAnswerContainer)

	err := quizService.InsertAnswers(userAnswerContainer)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
}

func getQuizResults(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)

	params := mux.Vars(r)

	userIDStr := params["id"]

	userID, err := strconv.Atoi(userIDStr)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	userReport, err := quizService.GetUserReport(userID)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(userReport)

}

var quizService QuizService
