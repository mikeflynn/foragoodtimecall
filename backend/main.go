package main

import (
	//"encoding/json"
	//"errors"
	//"net/http"
	"log"

	//"github.com/codegangsta/negroni"
	//"github.com/gorilla/context"
	//"github.com/gorilla/mux"
)

func main() {
	MySQLInit()
	defer MySQLClose()

	//MySQLInsert("contest", map[string]interface{}{"title": "A test"})
	//MySQLUpdate("contest", map[string]interface{}{"title": "An updated test"}, map[string]interface{}{"id": 1})
	//log.Println(MySQLSelect("contest", map[string]interface{}{"id": []interface{}{">", "1"}}, []string{"title", "winner_id"}))
	log.Println(MySQLSelect("contest", map[string]interface{}{"id": 1}, []string{"title", "winner_id"}))

	/*
		router := mux.NewRouter()
		apiRouter := mux.NewRouter()

		     // Main Routes
		   	router.HandleFunc("/", app.Handler).Methods("GET")

		     // API Routes
		     apiRouter.HandleFunc("/api/test", func() {

		     })

		   	router.PathPrefix("/api/").Handler(negroni.New(
		   		negroni.HandlerFunc(validateRequest),
		   		negroni.HandlerFunc(verifyJSON),
		   		negroni.Wrap(apiRouter),
		   	))


		n := negroni.Classic()
		//n.Use(sessions.Sessions("my_session", store))
		n.UseHandler(router)
		n.Run(":8081")
	*/
}
