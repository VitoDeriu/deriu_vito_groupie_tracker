package routeur

import (
	"GroupieTrackerJJBA/controleur"
	"net/http"
)

func InitServer() {

	http.HandleFunc("/", controleur.Index)

}
