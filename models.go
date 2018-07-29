package main

//Question - the question sent to the user
type Question struct {
	ID      int               `json:"id,omitempty"`
	Text    string            `json:"text,omitempty"`
	Options map[string]Option `json:"options,omitempty"`
}

type Option struct {
	ID         int    `json:"id,omitempty"`
	Label      string `json:"label,omitempty"`
	Text       string `json:"text,omitempty"`
	Correct    bool   `json:"-"`
	QuestionID int    `json:"questionId,omitempty"`
}

type Answer struct {
	QuestionID     int      `json:"questionId,omitempty"`
	CorrectOptions []string `json:"correctOptions,omitempty"`
}

type UserAnswerContainer struct {
	UserID  int            `json:"userId,omitempty"`
	Answers map[int]string `json:"answers,omitempty"` //key : questionId  value : label
}

type UserAnswerResultContainer struct {
	UserID         int `json:"id,omitempty"`
	CorrectAnswers int `json:"correctAnswers"`
}

type UserReport struct {
	UserID         int    `json:"userId"`
	CorrectAnswers int    `json:"correctAnswers"`
	Percentage     string `json:"percentage"`
}
