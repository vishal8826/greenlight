package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vishal8826/greenlight/internal/data"
)

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	data := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Oppenheimer",
		Year:      2023,
		Runtime:   230,
		Genres:    []string{"Sci-fi", "history"},
		Version:   1,
	}

	movie := envelope{"movie": data}

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logError(r, err)
		app.serverErrorResponse(w, r)
	}

}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}
