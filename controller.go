package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Controller struct {
	store  Storage
	mailer Mailer
}

func NewController(store Storage, mailer Mailer) *Controller {
	return &Controller{store: store, mailer: mailer}
}

func (s *Controller) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/users/{userID}/parse-kindle-file", s.handleParseKindleFile).Methods("POST")
	router.HandleFunc("/cloud/send-daily-insights", s.handleSendDailyInsights).Methods("POST")
}

func (s *Controller) handleParseKindleFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	file, _, err := r.FormFile("file")
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}

	defer file.Close()

	// parse that multipart file
	raw, err := parseKindleExtractFile(file)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, fmt.Sprintf("Error parsing file: %v", err))
		return
	}

	userIDint, _ := strconv.Atoi(userID)
	if err := s.createDataFromRawBook(raw, userIDint); err != nil {
		WriteJSON(w, http.StatusInternalServerError, fmt.Sprintf("Error creating data from raw book: %v", err))
		return
	}

	WriteJSON(w, http.StatusOK, "Successfully parsed file")
}

func (s *Controller) handleSendDailyInsights(w http.ResponseWriter, r *http.Request) {
	// get users
	users, err := s.store.GetUsers()
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// loop over users and get random highlights
	for _, u := range users {
		hs, err := s.store.GetRandomHighlights(3, u.ID)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		if len(hs) == 0 {
			continue
		}

		insights, err := s.buildInsights(hs)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = s.mailer.SendInsights(insights, u)
		if err != nil {
			WriteJSON(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	WriteJSON(w, http.StatusOK, nil)
}

// parseKindleExtractFile decodes a JSON file into a RawExtractBook struct.
func parseKindleExtractFile(file multipart.File) (*RawExtractBook, error) {
	decoder := json.NewDecoder(file)

	var raw RawExtractBook
	if err := decoder.Decode(&raw); err != nil {
		return nil, err
	}

	return &raw, nil
}

func (s *Controller) createDataFromRawBook(raw *RawExtractBook, userID int) error {
	_, err := s.store.GetBookByISBN(raw.ASIN)
	if err != nil {
		s.store.CreateBook(Book{
			ISBN:    raw.ASIN,
			Title:   raw.Title,
			Authors: raw.Authors,
		})
	}

	// create highlights
	hs := make([]Highlight, len(raw.Highlights))
	for i, h := range raw.Highlights {
		hs[i] = Highlight{
			Text:     h.Text,
			Location: h.Location.URL,
			Note:     h.Note,
			UserID:   userID,
			BookID:   raw.ASIN,
		}
	}

	err = s.store.CreateHighlights(hs)
	if err != nil {
		log.Println("Error creating highlights: ", err)
		return err
	}

	return nil
}

func (s *Controller) buildInsights(hs []*Highlight) ([]*DailyInsight, error) {
	var insights []*DailyInsight

	for _, h := range hs {
		book, err := s.store.GetBookByISBN(h.BookID)
		if err != nil {
			return nil, err
		}

		insights = append(insights, &DailyInsight{
			Text:        h.Text,
			Note:        h.Note,
			BookAuthors: book.Authors,
			BookTitle:   book.Title,
		})
	}

	return insights, nil
}
