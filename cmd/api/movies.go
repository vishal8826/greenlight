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
		http.NotFound(w, r)
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
		app.logger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}
