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
	router.HandleFunc("/answers", insertAnswers).Methods("POST")
	router.HandleFunc("/user/{id}/results", getQuizResults).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getQuestions(w http.ResponseWriter, r *http.Request) {

	questions, err := quizService.GetQuestions()

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	json.NewEncoder(w).Encode(questions)
}

func insertAnswers(w http.ResponseWriter, r *http.Request) {
	var userAnswerContainer *UserAnswerContainer

	_ = json.NewDecoder(r.Body).Decode(&userAnswerContainer)

	err := quizService.InsertAnswers(userAnswerContainer)

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
}

func getQuizResults(w http.ResponseWriter, r *http.Request) {
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
