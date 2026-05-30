package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	db "task-prioritizer/internal/database"
	"task-prioritizer/internal/models"
)

func Welcome(c *gin.Context) {
	if _, err := c.Cookie("user_id"); err != nil {
		c.HTML(200, "welcome.html", nil)
		return
	}
	c.Redirect(303, "/tasks")
}

func CreateTask(c *gin.Context) {
	type CreateTaskInput struct {
	Title          string `form:"title"`
	Importance     int    `form:"importance"`
	Difficulty     int    `form:"difficulty"`
	EstimatedHours int    `form:"estimated_hours"`
	Deadline       time.Time `form:"deadline" time_format:"2006-01-02"`
	}
	var input CreateTaskInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
	})
		return
	}
	userID, err := c.Cookie("user_id")
	if err != nil {
		c.String(500, "Зарегайся сначала")
		return
	}
	user_id, err := strconv.Atoi(userID)
	if err != nil {
		c.String(500, "гг")
		return
	}
	task := models.Task{
		Title: input.Title,
		Deadline: input.Deadline,
		Importance: input.Importance,
		Difficulty: input.Difficulty,
		EstimatedHours: input.EstimatedHours,
		UserID: user_id,
	}
	task.CalculatePriority()
	_, err = db.DB.Exec(context.Background(),
	`INSERT INTO tasks
	(title, deadline, importance, difficulty, estimated_hours, priority, user_id)
	VALUES ($1,$2,$3,$4,$5,$6, $7)`,
	task.Title,
	task.Deadline,
	task.Importance,
	task.Difficulty,
	task.EstimatedHours,
	task.Priority,
	task.UserID,
	)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/tasks")
}

func ListTasks(c *gin.Context) {
	userID, err := c.Cookie("user_id")
	if err != nil {
		c.Redirect(303, "/")
		return
	}
	user_id, err := strconv.Atoi(userID)
	if err != nil {
		c.String(500, "гг")
		return
	}
	rows, _ := db.DB.Query(context.Background(),
		`SELECT id,title,priority,deadline,difficulty,estimated_hours,importance FROM tasks WHERE user_id = $1
		 ORDER BY priority DESC`, user_id)
	type TaskView struct {
		ID             int
		Title          string
		Deadline       string
		Importance     int
		Difficulty     int
		EstimatedHours int
		Priority       float64
	}
	var tasks []TaskView
	for rows.Next() {
		var t models.Task
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Priority,
			&t.Deadline,
			&t.Difficulty,
			&t.EstimatedHours,
			&t.Importance,
		)
		if err != nil {
			c.String(500, err.Error())
			return
		}
		tasks = append(tasks, TaskView{
			ID: t.ID,
			Title: t.Title,
			Deadline: t.Deadline.Format("2006-01-02"),
			Importance: t.Importance,
			Difficulty: t.Difficulty,
			EstimatedHours: t.EstimatedHours,
		})
	}
	c.HTML(200, "tasks.html", gin.H{
		"tasks": tasks,
	})
}

func DeleteTask(c *gin.Context) {
	userID, err := c.Cookie("user_id")
    if err != nil {
        c.Redirect(303, "/")
        return
    }
    userIDInt, err := strconv.Atoi(userID)
    if err != nil {
        c.String(500, err.Error())
        return
    }
    taskID := c.Param("id")
    _, err = db.DB.Exec(
        context.Background(),
        `DELETE FROM tasks
         WHERE id = $1 AND user_id = $2`,
        taskID,
        userIDInt,
    )
    if err != nil {
        c.String(500, err.Error())
        return
    }
    c.Redirect(http.StatusSeeOther, "/tasks")
}