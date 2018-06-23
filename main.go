package main

import (
	"fmt"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/nerijusza/go-notes/pkg/render"
	"github.com/nerijusza/go-notes/pkg/storage"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	storageFactory := storage.Factory{}

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		storager, err := storageFactory.Get()
		display(ctx, storager, err, "")
	})

	// Method:   GET
	// Resource: http://localhost:8080/delete/{ID}
	app.Handle("GET", "/delete/{ID}", func(ctx iris.Context) {
		storager, err := storageFactory.Get()
		message := ""

		if err == nil {
			ID, err := ctx.Params().GetInt("ID")
			if err == nil {
				err = storager.Delete(ID)
				if err == nil {
					message = fmt.Sprintf("Message with id %v was deleted.", ID)
				}
			}
		}

		display(ctx, storager, err, message)
	})

	// Method:   POST
	// Resource: http://localhost:8080/add
	app.Handle("POST", "/add", func(ctx iris.Context) {
		storager, err := storageFactory.Get()
		noteContent := ctx.PostValue("note_content")
		message := ""

		if err == nil && noteContent != "" {
			noteID, err := storager.Add(noteContent)
			if err == nil {
				message = fmt.Sprintf("Message with ID %v was created!", noteID)
			}
		}

		display(ctx, storager, err, message)
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func display(ctx iris.Context, storager storage.Storager, err error, message string) {
	var notes = make([]storage.Note, 0)
	if err == nil {
		notes, err = storager.Get()
	}

	data := render.Data{notes, "", message}
	if err != nil {
		data.Error = err.Error()
	}

	render := render.Render{}
	renderedContent, renderError := render.Process(data)

	if renderError != nil {
		ctx.HTML("Render error: " + renderError.Error())
	} else {
		ctx.HTML(string(renderedContent))
	}
}
