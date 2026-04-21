package models

import (
	"errors"
	"time"
)

type Task struct {
	ID             int
	Title          string
	Deadline       time.Time
	Importance     int
	Difficulty     int
	EstimatedHours int
	Priority       float64
}

func (t *Task) Validate() error {

	if t.Title == "" {
		return errors.New("Название не может быть пустым")
	}

	if t.Importance < 1 || t.Importance > 5 {
		return errors.New("Важность должна быть от 1 до 5")
	}

	if t.Difficulty < 1 || t.Difficulty > 5 {
		return errors.New("Сложность должна быть от 1 до 5")
	}

	return nil
}

func (t *Task) CalculatePriority() {
	daysLeft := t.Deadline.Sub(time.Now()).Hours() / 24
	if daysLeft < 0 {
		daysLeft = 0
	}
	deadlineScore := 10 / (daysLeft + 1)
	importanceScore := float64(t.Importance * 2)
	difficultyScore := float64(t.Difficulty)
	timeScore := float64(t.EstimatedHours) / 2

t.Priority = deadlineScore + importanceScore + difficultyScore + timeScore
}