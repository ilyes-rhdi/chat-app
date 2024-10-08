package main

import (
	"fmt"
	"net/http"
	"Devenir_dev/cmd/handlers"
	"log"
	"github.com/gorilla/mux"
)

const port = ":3000"
func main (){
	handlers.InitDB()
	app := mux.NewRouter()
	app.HandleFunc("/login", handlers.Login)
	app.HandleFunc("/Submit", handlers.Submit)
	app.HandleFunc("/Home", handlers.Main)
	app.HandleFunc("/deleteUser", handlers.DeleteUserHandler)
	fmt.Println("(http://localhost:3000/login) le serveur est lancer sur ce lien ")
	
	err := http.ListenAndServe(port, app)
	if err != nil {
		log.Fatal(err)
	}
}
  



 