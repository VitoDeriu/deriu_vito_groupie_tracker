package router

import (
	"GroupieTrackerJJBA/controller"
	"fmt"
	"net/http"
	"os"
)

func InitServer() {

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/allstands", controller.AllStands)
	http.HandleFunc("/allcharacters", controller.AllCharacters)
	http.HandleFunc("/character", controller.DisplayCharacter)
	http.HandleFunc("/stand", controller.DisplayStand)
	http.HandleFunc("/recherche", controller.DisplayRecherche)
	http.HandleFunc("/recherchetreatment", controller.Recherche)
	http.HandleFunc("/favoris", controller.Favoris)
	http.HandleFunc("/favoritreatment", controller.FavoriTreatment)
	http.HandleFunc("/filtretreatment", controller.Filtre)
	http.HandleFunc("/favoridelete", controller.FavoriDelete)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("(http://localhost:8080/) - Server started on port:8080")
	http.ListenAndServe("localhost:8080", nil)

}
