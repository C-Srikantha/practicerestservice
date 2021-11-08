package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/gorilla/mux"
	"rest.com/service/databasecon"
	"rest.com/service/handler"
)

type Actor struct {
	Id         uint   `pg:",pk"`
	ActorName  string `pg:",notnull"`
	ActorPhone int64  `pg:",notnull"`
}
type Movie struct {
	Id        uint   `pg:",pk"`
	MovieName string `pg:",notnull"`
	MovieLang string `pg:",notnull"`
	MovieType string `pg:",notnull"`
	ActorID   uint
	Actor     *Actor `pg:"rel:has-many"`
}
type Movierelease struct {
	Releaseyear int
	MovieId     uint
	Movie       *Movie `pg:"rel:has-one"`
}

//creates a table in database postpresql
func createtable(db *pg.DB) {
	model := []interface{}{
		(*Actor)(nil),
		(*Movie)(nil),
		(*Movierelease)(nil),
	}
	for _, model := range model {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			Varchar:       50,
			FKConstraints: true,
		})

		if err != nil {
			fmt.Println(err)
		}
	}
}

//handles the request and creates the server
func handelrequest(db *pg.DB) {
	mux := mux.NewRouter()
	mux.HandleFunc("/getdetails", func(w http.ResponseWriter, r *http.Request) { handler.Getdetails(w, r, db) }).Methods("GET")               //registers api signature and handler to the router
	mux.HandleFunc("/getdetails/{id}", func(w http.ResponseWriter, r *http.Request) { handler.Getadetails(w, r, db) }).Methods("GET")         //registers api signature and handler to the router
	mux.HandleFunc("/postdetails", func(w http.ResponseWriter, r *http.Request) { handler.Postdetails(w, r, db) }).Methods("POST")            //registers api signature and handler to the router
	mux.HandleFunc("/deletedetails/{id}", func(w http.ResponseWriter, r *http.Request) { handler.Deletedetails(w, r, db) }).Methods("DELETE") //registers api signature and handler to the router
	mux.HandleFunc("/updatedetails/{id}", func(w http.ResponseWriter, r *http.Request) { handler.Updatedetails(w, r, db) }).Methods("PUT")    //registers api signature and handler to the router
	mux.HandleFunc("/getdetails/movie", func(w http.ResponseWriter, r *http.Request) { handler.Getmoviedetails(w, r, db) }).Methods("GET")    //registers api signature and handler to the router
	mux.HandleFunc("/getdetails/movie/{id}", func(w http.ResponseWriter, r *http.Request) { handler.Getamoviedetails(w, r, db) }).Methods("GET")
	mux.HandleFunc("/postdetails/movie", func(w http.ResponseWriter, r *http.Request) { handler.Postcastdetails(w, r, db) }).Methods("POST")
	mux.HandleFunc("/updatedetails/movie/{id}", func(w http.ResponseWriter, r *http.Request) { handler.Updatecastdetails(w, r, db) }).Methods("PUT")
	//mux.HandleFunc("/getapi", func(w http.ResponseWriter, r *http.Request) { handler.Getapidetails(w, r, db) }).Methods("GET")
	fmt.Println(http.ListenAndServe(":8081", mux)) //creates a server and listens to the port for requests and pass the requests to route
}
func main() {
	db, err := databasecon.Setup() //calling database connection
	defer db.Close()
	if db == nil || err != nil {
		fmt.Println("Connection failed!!")
		os.Exit(1) //terminates the server
	}
	createtable(db)

	handler.Readfile(db)
	handelrequest(db)

}
