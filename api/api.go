package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/patrickhener/go-cj/config"
)

type Session struct {
	ID                  int     `json:"id"`
	Description         string  `json:"description"`
	Name                string  `json:"name"`
	Username            string  `json:"username"`
	TerminateAt         string  `json:"terminateAt"`
	Active              bool    `json:"active"`
	NotificationEnabled bool    `json:"notificationEnabled"`
	CreatedAt           string  `json:"createdAt"`
	Hashcat             Hashcat `json:"hashcat"`
}

type Hashcat struct {
	State            int    `json:"state"`
	CrackedPasswords int    `json:"crackedPasswords"`
	AllPasswords     int    `json:"allPasswords"`
	Progress         string `json:"progress"`
	TimeRemaining    string `json:"timeRemaining"`
	ETA              string `json:"estimatedCompletionTime"`
}

func sendRequest(endpoint, method string) ([]byte, error) {
	uri := fmt.Sprintf("%s%s", config.GoCJConfig.BaseURI, endpoint)
	client := &http.Client{}

	// build request
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}

	// add header
	req.Header.Add("X-CrackerJack-Auth", config.GoCJConfig.Auth)

	// send
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// defer resp.Body.Close()

	// check status
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Request was not successful: %s", resp.Status)
	}
	// return result body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func getAllSessions() (string, error) {
	var sessions []Session

	result, err := sendRequest("sessions", "GET")
	if err != nil {
		return "", err
	}

	// parse results
	if err := json.Unmarshal(result, &sessions); err != nil {
		return "", err
	}

	// Render table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "Description", "Username", "Active"})
	for _, s := range sessions {
		table.Append([]string{fmt.Sprint(s.ID), s.Description, s.Username, fmt.Sprint(s.Active)})
	}
	table.Render()

	return "", nil
}

func getSpecificSession(id int) (string, error) {
	var s Session

	result, err := sendRequest(fmt.Sprintf("sessions/%d", id), "GET")
	if err != nil {
		return "", err
	}

	// parse results
	if err := json.Unmarshal(result, &s); err != nil {
		return "", err
	}

	fmt.Printf("\nID: %d\nDescription: %s\nUsername: %s\nActive: %+v\nCreated At: %s\nTerminate At: %s\nNotifications: %+v\nHashes provided: %d\nHashes cracked: %d\nProgress: %s %%\nTime Remaining: %s\nCompleted At: %s\n\n", s.ID, s.Description, s.Username, s.Active, s.CreatedAt, s.TerminateAt, s.NotificationEnabled, s.Hashcat.AllPasswords, s.Hashcat.CrackedPasswords, s.Hashcat.Progress, s.Hashcat.TimeRemaining, s.Hashcat.ETA)

	return "", nil
}
