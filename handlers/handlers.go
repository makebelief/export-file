package handler

import (
	"html/template"
	"net/http"
	"strings"
	"strconv"

	"piscine/ascii"
)

type Data struct {
	Text   string
	Result string
}

// RenderErrorPage renders an error page based on the status code
func RenderErrorPage(w http.ResponseWriter, statusCode int) {
	var templatePath string

	switch statusCode {
	case http.StatusNotFound:
		templatePath = "templates/404.html"
	case http.StatusInternalServerError:
		templatePath = "templates/500.html"
	case http.StatusMethodNotAllowed:
		templatePath = "templates/405.html"
	case http.StatusBadRequest:
		templatePath = "templates/400.html"
	default:
		templatePath = "templates/500.html" // Default to 500 for unknown status codes
	}

	w.WriteHeader(statusCode)
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	RenderErrorPage(w, http.StatusNotFound)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderErrorPage(w, http.StatusNotFound)
		return
	}

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			RenderErrorPage(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, nil)
		if err != nil {
			RenderErrorPage(w, http.StatusInternalServerError)
		}
	} else {
		RenderErrorPage(w, http.StatusMethodNotAllowed)
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		font := r.FormValue("file")

		for _, char := range text {
			if char > '~' {
				RenderErrorPage(w, http.StatusBadRequest)
				return
			}
		}

		if text == "" || font == "" {
			RenderErrorPage(w, http.StatusBadRequest)
			return
		}

		var print string

		if strings.Contains(text, "\n") {
			myslice := strings.Split(text, "\n")
			for _, line := range myslice {
				print += ascii.PrintAsci(font, line)
			}
		} else {
			print = ascii.PrintAsci(font, text)
			if print == "nil" {
				RenderErrorPage(w, http.StatusNotFound)
				return
			}
		}

		data := Data{
			Text:   text,
			Result: print,
		}
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			RenderErrorPage(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			RenderErrorPage(w, http.StatusInternalServerError)
		}
	} else {
		RenderErrorPage(w, http.StatusMethodNotAllowed)
	}
}

// ExportAsciiArtHandler handles the export of ASCII art
func ExportAsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		RenderErrorPage(w, http.StatusMethodNotAllowed)
		return
	}

	// Get the ASCII art result from query parameters
	result := r.URL.Query().Get("result")
	if result == "" {
		RenderErrorPage(w, http.StatusBadRequest)
		return
	}

	// Set the headers for file download
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
	w.Header().Set("Content-Length", strconv.Itoa(len(result)))

	// Write the ASCII art to the response
	if _, err := w.Write([]byte(result)); err != nil {
		RenderErrorPage(w, http.StatusInternalServerError)
	}
}
