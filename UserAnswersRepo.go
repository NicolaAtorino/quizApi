package main

import "errors"

type UserAnswersRepo struct {
}

var containerList = make(map[int]*UserAnswerResultContainer)

func (s *UserAnswersRepo) InsertUserAnswerResults(container *UserAnswerResultContainer) error {

	containerList[container.UserID] = container
	return nil
}
func (s *UserAnswersRepo) GetTotalUsers() int {
	return len(containerList)
}

func (s *UserAnswersRepo) GetUserAnswerResults(userID int) (*UserAnswerResultContainer, error) {

	container, containerFound := containerList[userID]

	if !containerFound {
		return nil, errors.New("user results not found")
	}
	return container, nil
}
