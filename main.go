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

type Movie struct {
	Id        uint16 `pg:",pk"`
	MovieName string `pg:",notnull"`
	MovieLang string `pg:",notnull"`
	MovieType string `pg:",notnull"`
}
type Cast struct {
	Id           uint16 `pg:",fk:"`
	ProducerName string `pg:",notnull"`
	DirectorName string `pg:",notnull"`
	ActorName    string `pg:",notnull"`
}

//creates a table in database postpresql
func createtable(db *pg.DB) {
	model := []interface{}{
		(*Movie)(nil),
		(*Cast)(nil),
	}
	for _, model := range model {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			Varchar:     50,
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
	mux.HandleFunc("/getdetails/cast", func(w http.ResponseWriter, r *http.Request) { handler.Getcastdetails(w, r, db) }).Methods("GET")      //registers api signature and handler to the router
	mux.HandleFunc("/getdetails/cast/{id}", func(w http.ResponseWriter, r *http.Request) { handler.Getacastdetails(w, r, db) }).Methods("GET")
	mux.HandleFunc("/postdetails/cast", func(w http.ResponseWriter, r *http.Request) { handler.Postcastdetails(w, r, db) }).Methods("POST")
	mux.HandleFunc("/updatedetails/cast/{id}", func(w http.ResponseWriter, r *http.Request) { handler.Updatecastdetails(w, r, db) }).Methods("PUT")

	fmt.Println(http.ListenAndServe(":8081", mux)) //creates a server and listens to the port for requests and pass the requests to route
}
func main() {
	db := databasecon.Setup() //calling database connection
	defer db.Close()
	if db == nil {
		fmt.Println("Connection failed!!")
		os.Exit(1) //terminates the server
	}
	createtable(db)
	handelrequest(db)
}
