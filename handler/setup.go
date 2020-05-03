package handler

import (
	"context"
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/gcal-timetracker/alfred"
	"github.com/hayashiki/gcal-timetracker/calendar"
)

func DoSetup(wf *aw.Workflow, _ []string) (string, error) {
	token, err := alfred.GetToken(wf)
	if err != nil {
		return "", fmt.Errorf("please authorize")
	}

	clientID := alfred.GetClientID(wf)
	client, err := calendar.NewClient(context.Background(), calendar.NewConfig(clientID), token)

	if err != nil {
		return "", fmt.Errorf("something wrong happened, please try again later 🙏 (%w)", err)
	}

	id, err := client.CreateCalendar()
	if err != nil {
		return "", fmt.Errorf("could not create the calendar, please try again later 🙏 (%w)", err)
	}

	if err := alfred.SetCalendarID(wf, id); err != nil {
		return "", fmt.Errorf("cannot save the configuration in Alfred, please try again later 🙏 (%w)", err)
	}

	return "Calendar created successfully 📅", nil
}
