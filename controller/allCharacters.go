package controller

import (
	"GroupieTrackerJJBA/data"
	"GroupieTrackerJJBA/temps"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var CharacterList []data.Character
var PageChar int

func AllCharacters(w http.ResponseWriter, r *http.Request) {

	//struct globale pour l'envoie des données dans le template
	var Data data.DataPaginate

	//recup du numero de page dans la query
	PageChar, _ = strconv.Atoi(r.URL.Query().Get("page"))

	//call api de tous les characters
	CallAllChar()

	//pagination puis stockage des 10 characters dans la struct globale
	characterListPaginated, paginationData := PaginateCharacter(CharacterList)
	Data.CharList = characterListPaginated
	Data.PageData = paginationData

	temps.Temp.ExecuteTemplate(w, "AllCharacters", &Data)
}

// fonction de pagination pour la page des Characters.
// on recup la liste des charracters pour renvoyer une autre liste avec seulement 10 characters pour les afficher dans le template
// on envoi également la struct de pagination pour afficher le numero de la page et creer les liens Next et Previous page
func PaginateCharacter(listToPaginate []data.Character) ([]data.Character, data.Page) {

	var Pagination data.Page

	perPage := 8                                  //nombre d'elements par page
	offSet := 8                                   //decalage pour l'affichage grace aux index
	nbPage := (len(listToPaginate) / perPage)     //nombre de page
	nbLastPage := (len(listToPaginate) % perPage) //nombre d'éléments dans la dernière page (doit être inférieur a nbPage)

	//assignation des numero de page par rapport a la currentPage
	Pagination.PreviousPage = PageChar - 1
	Pagination.CurrentPage = PageChar
	Pagination.NextPage = PageChar + 1
	Pagination.MaxPage = nbPage

	//changement du offset pour la derniere page lorsqu'il y a moins de 10 elements sur la dernière page.
	if Pagination.CurrentPage == Pagination.MaxPage {
		offSet = nbLastPage
	}

	//stockage de la nouvelle liste de 10 elements dans displaychar
	Displaychar := listToPaginate[(Pagination.CurrentPage * perPage):(Pagination.CurrentPage*perPage + offSet)]

	//return de la liste de chararcters et de la struct de pagination
	return Displaychar, Pagination
}

// appel et marshal le json de tous les characters et renvoi un tableau de struct
func CallAllChar() {
	println("call de all characters")
	urlApi := "https://stand-by-me.herokuapp.com/api/v1/characters/"
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("erreur lors de la requete Api", errReq.Error())
	}

	req.Header.Set("Content-Type", "application/json")

	res, errRes := httpClient.Do(req)
	if res.Body != nil {
		defer res.Body.Close()
	} else {
		fmt.Println("erreur lors de l'envoie de la requete", errRes.Error())
	}

	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("erreur lors de la lecture du corps de la requete", errBody.Error())
	}

	json.Unmarshal(body, &CharacterList)
	println("call done")
}
