package main

import (
	"log"
	"net/http"
)

func main() {
	// 1
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// 2
	// router := mux.NewRouter()
	// router.HandleFunc("/", Index)
	// log.Fatal(http.ListenAndServe(":8080", router))

	// 3
	// router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/", Index)
	// router.HandleFunc("/todos", TodoIndex)
	// router.HandleFunc("/todos/{todoID}", TodoShow)

	// log.Fatal(http.ListenAndServe(":8080", router))

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":9999", router))

}

// 2
// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
// }

// 3
// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Welcome")
// }

// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	todos := Todos{
// 		Todo{
// 			Name: "Write presentation",
// 		},
// 		Todo{
// 			Name: "Host meetup",
// 		},
// 	}
// 	json.NewEncoder(w).Encode(todos)
// 	// fmt.Fprintln(w, "Todo Index!")
// }

// func TodoShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	todoID := vars["todoID"]
// 	fmt.Fprintln(w, "Todo show:", todoID)
// }
