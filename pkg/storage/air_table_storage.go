package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nerijusza/go-notes/pkg/helper"
)

type AirTableFullNote struct {
	ID          string
	Fields      Note
	CreatedTime string
}

type AirTableListRecordsResponse struct {
	Records []AirTableFullNote
}

type AirTableCreateRecordRequest struct {
	Fields Note `json:"fields"`
}

// AirTableStorage implementation of StorageInterface in airtable.com
type AirTableStorage struct {
	Account   string
	APIKey    string
	TableName string
}

// Get gets all notes sorted fron newest to oldest
func (t *AirTableStorage) Get() ([]Note, error) {
	airTableFullNotes, err := t.getAirTaBleNotes()

	notes := make([]Note, 0)
	for _, airTableNote := range airTableFullNotes {
		notes = append(notes, airTableNote.Fields)
	}

	fmt.Println(notes)
	fmt.Println(err)

	return notes, err
}

// GetN get newest n notes
func (t *AirTableStorage) GetN(quantity int) ([]Note, error) {
	notes, err := t.Get()
	if err != nil {
		return nil, err
	}

	if len(notes) > quantity {
		notes = notes[0:quantity]
	}

	return notes, nil
}

// Add saves given string as note
func (t *AirTableStorage) Add(content string) (int, error) {
	ID, err := t.getBiggestID()
	if err != nil {
		return 0, err
	}

	data := AirTableCreateRecordRequest{Note{ID + 1, content, helper.GetCurrentTime()}}
	json, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	url := t.getMainPath() + t.Account + "/" + t.TableName

	client := &http.Client{}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	request.Header.Add("Authorization", `Bearer `+t.APIKey)
	request.Header.Add("Content-type", `application/json`)
	response, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	if response.StatusCode != http.StatusOK {
		return 0, t.formatFriendlyError(response)
	}

	return data.Fields.ID, nil
}

// Delete deletes given note by id
func (t *AirTableStorage) Delete(ID int) error {
	airTableNotes, err := t.getAirTaBleNotes()
	if err != nil {
		return nil
	}

	for _, airTableNote := range airTableNotes {
		if airTableNote.Fields.ID == ID {
			return t.deleteByAirTableID(airTableNote.ID)
		}
	}

	return nil
}

// DeleteAll deletes all notes from storage
func (t *AirTableStorage) DeleteAll() error {
	airTableNotes, err := t.getAirTaBleNotes()
	if err != nil {
		return err
	}

	for _, airTableNote := range airTableNotes {
		err = t.deleteByAirTableID(airTableNote.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// getBiggestID return biggest note ID in database
func (t *AirTableStorage) getBiggestID() (int, error) {
	notes, err := t.Get()
	if err != nil {
		return 0, err
	}

	if len(notes) > 0 {
		return notes[0].ID, nil
	}

	return 0, nil
}

func (t AirTableStorage) getMainPath() string {
	return "https://api.airtable.com/v0/"
}

// GetAirTaBleNotes gets all notes (Air Table structure with internal id) sorted fron newest to oldest
func (t *AirTableStorage) getAirTaBleNotes() ([]AirTableFullNote, error) {
	url := t.getMainPath() + t.Account + "/" + t.TableName + "?sort%5B0%5D%5Bfield%5D=ID&sort%5B0%5D%5Bdirection%5D=desc&api_key=" + t.APIKey

	res, err := http.Get(url)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := AirTableListRecordsResponse{}
	err = json.Unmarshal(body, &response)

	return response.Records, err
}

func (t *AirTableStorage) deleteByAirTableID(ID string) error {
	url := t.getMainPath() + t.Account + "/" + t.TableName + "/" + ID

	client := &http.Client{}
	request, err := http.NewRequest("DELETE", url, nil)
	request.Header.Add("Authorization", `Bearer `+t.APIKey)
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return t.formatFriendlyError(response)
	}

	return nil
}

func (t AirTableStorage) formatFriendlyError(response *http.Response) error {
	content, _ := ioutil.ReadAll(response.Body)
	return fmt.Errorf("Air table server returned code: %v, error: %v", response.StatusCode, fmt.Sprintf("%s", content))
}
