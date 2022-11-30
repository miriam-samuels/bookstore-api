 package routes
 
 import ( 
	 "github.com/gorilla/mux" 
	  "github.com/miriam-samuels/bookstore/pkg/controllers" 
	)
	// gorilla mux for dynamic routing instead of  http.HandleFunc for static
	
	var BookStoreRoutes = func (r *mux.Router){
		r.handleFunc("/book/",controllers.CreateBook).Methods("POST")
		r.handleFunc("/book/",controllers.GetBooks).Methods("GET")
		r.handleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
		r.handleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
		r.handleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
	}