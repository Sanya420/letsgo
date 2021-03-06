package main

import (
	"awesomeProject15/pkg/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home_page.html", &templateData{
		Snippets: s,
	})

	//	files := []string{
	//		"./ui/html/home_page.html",
	//		"./ui/html/base_layout.html",
	//		"./ui/html/footer_partial.html",
	//}

	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
	//err = ts.Execute(w, data)
	//if err != nil {
	//	app.serverError(w, err)
	//}
}

func (app *application) secondpage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	app.render(w, r, "show_page.html", &templateData{
		Snippet: s,
	})

	//files := []string{
	//"./ui/html/show_page.html",
	//"./ui/html/base_layout.html",
	//"./ui/html/footer_partial.html",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//app.serverError(w, err)
	//return
	//}
	//
	//err = ts.Execute(w, data)
	//if err != nil {
	//	app.serverError(w, err)
	//}
}

func (app *application) raptext(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")

		app.clientError(w, http.StatusMethodNotAllowed)

		return
	}

	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/sec?id=%d", id), http.StatusSeeOther)
}
