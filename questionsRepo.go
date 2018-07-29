package main

import "errors"

type QuestionsRepo struct {
	Questions map[int]*Question
}

//var questions = make(map[int]*Question)

func (s *QuestionsRepo) GetQuestions() (map[int]*Question, error) {

	return s.Questions, nil
}

func (s *QuestionsRepo) FindQuestionByID(ID int) (*Question, error) {

	var question = s.Questions[ID]

	return question, nil
}

func (s *QuestionsRepo) GetTotalQuestionsNumber() int {
	return len(s.Questions)
}

func (s *QuestionsRepo) FindAnswerByQuestionId(ID int) (*Answer, error) {

	var question, questionFound = s.Questions[ID]

	if !questionFound {
		return nil, errors.New("question not found")
	}

	var correctOptions []string

	for lbl, opt := range question.Options {
		if opt.Correct {
			correctOptions = append(correctOptions, lbl)
		}
	}

	var answer = &Answer{
		QuestionID:     question.ID,
		CorrectOptions: correctOptions,
	}

	return answer, nil
}

func (s *QuestionsRepo) Init() {

	s.Questions = make(map[int]*Question)

	optionsQ1 := make(map[string]Option)
	optionsQ1["A"] = Option{
		ID:         11,
		QuestionID: 1,
		Label:      "A",
		Text:       "Question 1 - Answer 1 (correct)",
		Correct:    true,
	}
	optionsQ1["B"] = Option{
		ID:         12,
		QuestionID: 1,
		Label:      "B",
		Text:       "Question 1 - Answer 2",
		Correct:    false,
	}
	optionsQ1["C"] = Option{
		ID:         13,
		QuestionID: 1,
		Label:      "C",
		Text:       "Question 1 - Answer 3",
		Correct:    false,
	}
	optionsQ1["D"] = Option{
		ID:         14,
		QuestionID: 1,
		Label:      "D",
		Text:       "Question 1 - Answer 4",
		Correct:    false,
	}

	question1 := &Question{
		ID:      1,
		Text:    "Question 1",
		Options: optionsQ1,
	}

	s.Questions[question1.ID] = question1

	optionsQ2 := make(map[string]Option)
	optionsQ2["A"] = Option{
		ID:         21,
		QuestionID: 2,
		Label:      "A",
		Text:       "Question 2 - Answer 1",
		Correct:    false,
	}
	optionsQ2["B"] = Option{
		ID:         22,
		QuestionID: 2,
		Label:      "B",
		Text:       "Question 2 - Answer 2 (correct)",
		Correct:    true,
	}
	optionsQ2["C"] = Option{
		ID:         23,
		QuestionID: 2,
		Label:      "C",
		Text:       "Question 2 - Answer 3",
		Correct:    false,
	}
	optionsQ2["D"] = Option{
		ID:         24,
		QuestionID: 2,
		Label:      "D",
		Text:       "Question 2 - Answer 4",
		Correct:    false,
	}

	question2 := &Question{
		ID:      2,
		Text:    "Question 2",
		Options: optionsQ2,
	}

	s.Questions[question2.ID] = question2
}
