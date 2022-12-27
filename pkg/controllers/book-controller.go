package controllers

import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/miriam-samuels/bookstore/pkg/utils"
	"github.com/miriam-samuels/bookstore/pkg/models"
)

var Book models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book , _ := models.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request){
	NewBook := &models.Book{}
	utils.ParseBody(r, NewBook)
	b:=  NewBook.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book , _ := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err :=strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	updateBook := &models.Book{}
	utils.ParseBody(r, NewBook)

	bookDetails, db := models.GetBookById(ID)
	if updateBook.Name !- ""{
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author !- ""{
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication !- ""{
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res,_ := json.Marshal(bookDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}