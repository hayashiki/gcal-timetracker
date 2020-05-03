package handler

import (
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/google/uuid"
	"github.com/hayashiki/gcal-timetracker/alfred"
	"time"
)

func DoStart(wf *aw.Workflow, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("please provide some input 👀")
	}

	task := alfred.Task{ID: uuid.New().String(), Description: args[0], Start: time.Now()}

	tasks, _ := alfred.LoadOngoingTasks(wf)
	tasks = append(tasks, task)

	if err := alfred.StoreOngoingTasks(wf, tasks); err != nil {
		return "", fmt.Errorf("cannot store the new task, please try again later 🙏 (%w)", err)
	}

	return "Task started, remember to stop it 😉", nil
}
