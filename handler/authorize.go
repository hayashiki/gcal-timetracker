package handler

import (
	"encoding/json"
	"fmt"
	aw "github.com/deanishe/awgo"
	"github.com/hayashiki/gcal-timetracker/alfred"
	"github.com/hayashiki/gcal-timetracker/calendar"
)

func DoAuthorize(wf *aw.Workflow, _ []string) (string, error) {
	config := calendar.NewConfig(alfred.GetOAuth(wf))

	token, err := calendar.GetToken(config)
	if err != nil {
		return "", fmt.Errorf("cannot get an access token 😢 (%w)", err)
	}

	b, err := json.Marshal(token)
	if err != nil {
		return "", fmt.Errorf("cannot serialize the token to JSON 😢 (%w)", err)
	}

	if err := alfred.SetToken(wf, string(b)); err != nil {
		return "", fmt.Errorf("cannot store the token in the keychain 😢 (%w)", err)
	}

	return "Token stored successfully 😎", nil
}
