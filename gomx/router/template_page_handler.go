package router

import (
	"html/template"
	"log"
	"net/http"

	"github.com/winstonco/gomx/config"
)

type PageData struct {
	Arg any
}

type TemplatePageHandler struct {
	template *template.Template
	data     PageData
}

func (tph *TemplatePageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tph.template.Execute(w, tph.data)
	if err != nil {
		internalErrorHandler(err)(w, r)
	}
}

func internalErrorHandler(err error) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		t := template.Must(template.ParseFiles(
			config.BaseTemplate,
			config.ReservedDir+"/500.gohtml",
		))
		err := t.Execute(w, PageData{
			Arg: err,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	return http.HandlerFunc(fn)
}

func notFoundHandler() http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		t := template.Must(template.ParseFiles(
			config.BaseTemplate,
			config.ReservedDir+"/404.gohtml",
		))
		err := t.Execute(w, nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	return http.HandlerFunc(fn)
}
