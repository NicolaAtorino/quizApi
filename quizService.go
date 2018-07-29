package main

import (
	"errors"
	"fmt"
)

type QuizService struct {
	QuestionsRepo   *QuestionsRepo
	UserAnswersRepo *UserAnswersRepo
}

func (q *QuizService) Init() {

	q.QuestionsRepo = &QuestionsRepo{}
	q.QuestionsRepo.Init()

	q.UserAnswersRepo = &UserAnswersRepo{}

}

func (q *QuizService) GetQuestions() (map[int]*Question, error) {

	return q.QuestionsRepo.GetQuestions()

}

func (q *QuizService) InsertAnswers(container *UserAnswerContainer) error {

	questions, err := q.QuestionsRepo.GetQuestions()

	if err != nil {
		return err
	}

	correctAnswers := 0

	for questionID, optionLabel := range container.Answers {

		question, questionFound := questions[questionID]

		if !questionFound {
			return errors.New("questionId not found")
		}

		option, optionFound := question.Options[optionLabel]

		if !optionFound {
			return errors.New("option not found")
		}

		if option.Correct {
			correctAnswers++
		}
	}

	userAnswersResultContainer := &UserAnswerResultContainer{
		UserID:         container.UserID,
		CorrectAnswers: correctAnswers,
	}

	return q.UserAnswersRepo.InsertUserAnswerResults(userAnswersResultContainer)
}

func (q *QuizService) GetUserReport(userID int) (*UserReport, error) {

	container, err := q.UserAnswersRepo.GetUserAnswerResults(userID)

	if err != nil {
		return nil, err
	}

	totalUser := q.UserAnswersRepo.GetTotalUsers()

	lessPerformantUsers := getLessPerformantUsersTotal(userID, container.CorrectAnswers)

	floatTotal := float64(totalUser)
	floatLess := float64(lessPerformantUsers)
	userRatio := floatLess * float64(100) / floatTotal

	percentage := fmt.Sprintf("%f%%", userRatio)

	userReport := &UserReport{
		UserID:         userID,
		CorrectAnswers: container.CorrectAnswers,
		Percentage:     percentage,
	}

	return userReport, nil

}

func getLessPerformantUsersTotal(userID int, correctAnswers int) int {
	worstResults := 0
	for uID, container := range containerList {
		if uID == userID {
			continue //we don't calculate the same user
		}

		if container.CorrectAnswers < correctAnswers {
			worstResults++
		}
	}
	return worstResults
}
