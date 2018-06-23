package render

import (
	"bytes"
	"html/template"

	"github.com/nerijusza/go-notes/pkg/storage"
)

// Render struct responsible for html rendering
type Render struct {
}

// Data required input to render page
type Data struct {
	Notes   []storage.Note
	Error   string
	Message string
}

type splitedData struct {
	Notes   [][]storage.Note
	Error   string
	Message string
}

// Process renders html page
func (t *Render) Process(data Data) (string, error) {
	preparedData := splitedData{
		splitNotes(data.Notes, 4),
		data.Error,
		data.Message,
	}

	tmpl := template.New("index.htm")
	tmpl, err := tmpl.ParseFiles("public/index.htm")

	renderedOutput := &bytes.Buffer{}
	err = tmpl.ExecuteTemplate(renderedOutput, "index.htm", preparedData)
	return renderedOutput.String(), err
}

func splitNotes(notes []storage.Note, countOnRow int) [][]storage.Note {
	output := make([][]storage.Note, 0)

	row := make([]storage.Note, 0)
	for _, note := range notes {
		row = append(row, note)
		if len(row) == countOnRow {
			output = append(output, row)
			row = make([]storage.Note, 0)
		}
	}

	if len(row) > 0 {
		output = append(output, row)
	}

	return output
}
