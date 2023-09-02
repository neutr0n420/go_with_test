
package main

import(
	"fmt"
	"log"
	"net/http"
	"strconv"	
)

func home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/"{
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello form Snippetbox"))
}




func showSnippet(w http.ResponseWriter, r *http.Request)  {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err!=nil || id <1{
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippit with the Id %d ...", id)
}





func createSnipet(w http.ResponseWriter, r *http.Request){

	if r.Method != "POST"{
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not allowed"))
		http.Error(w, "Method not found ", 405)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}





func main()  {
	mux:= http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet",showSnippet)
	mux.HandleFunc("/snippet/create",createSnipet)
	

	log.Println("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
	
}
