package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	booksdto "waysbook/dto/books"
	dto "waysbook/dto/results"
	"waysbook/models"
	"waysbook/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerBook struct {
	BookRepository repositories.BookRepository
}

func HandlerBook(BookRepository repositories.BookRepository) *handlerBook {
	return &handlerBook{BookRepository}
}

func (h *handlerBook) FindBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	books, err := h.BookRepository.FindBooks()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	for i, p := range books {
		books[i].File = os.Getenv("PATH_FILE") + p.File
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: books}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var book models.Book
	book, err := h.BookRepository.GetBookByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book.File = os.Getenv("PATH_FILE") + book.File

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: book}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerBook) AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	coverContext := r.Context().Value("dataFile")
	filepath := coverContext.(string)

	pdfContext := r.Context().Value("dataPDF")
	filename := pdfContext.(string)

	price, _ := strconv.Atoi(r.FormValue("price"))
	pages, _ := strconv.Atoi(r.FormValue("pages"))

	request := booksdto.AddBookRequest{
		Title:           r.FormValue("title"),
		PublicationDate: r.FormValue("publication_date"),
		Pages:           pages,
		ISBN:            r.FormValue("isbn"),
		Price:           price,
		About:           r.FormValue("about"),
		File:            filename,
		Cover:           filepath,
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "WaysBook"})

	if err != nil {
		fmt.Println(err.Error())
	}

	validation := validator.New()
	err = validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book := models.Book{
		Title:           request.Title,
		PublicationDate: request.PublicationDate,
		Pages:           request.Price,
		ISBN:            request.ISBN,
		Price:           request.Price,
		About:           request.About,
		File:            filename,
		Cover:           resp.SecureURL,
	}

	book, err = h.BookRepository.AddBook(book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book, err = h.BookRepository.GetBookByID(book.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	book, _ = h.BookRepository.GetBookByID(book.ID)
	book.File = os.Getenv("PATH_FILE") + book.File

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: "success", Data: book}
	json.NewEncoder(w).Encode(response)
}
