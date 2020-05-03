package alfred

import (
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/gcal-timetracker/calendar"
)

const (
	clientID   = "client_id"
	clientSecret   = "client_secret"
	calendarID = "calendar_id"
)

func GetClientID(wf *aw.Workflow) string {
	return wf.Config.GetString(clientID)
}

func GetOAuth(wf *aw.Workflow) *calendar.AuthorizeAPI {
	var config calendar.AuthorizeAPI
	config.ClientID = wf.Config.GetString(clientID)
	config.ClientSecret = wf.Config.GetString(clientSecret)

	return &config
}


func GetCalendarID(wf *aw.Workflow) string {
	return wf.Config.Get(calendarID)
}

func SetCalendarID(wf *aw.Workflow, id string) error {
	return wf.Config.Set(calendarID, id, false).Do()
}
