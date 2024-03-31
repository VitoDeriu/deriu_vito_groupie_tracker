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

var StandList []data.Stand
var PageStand int

func AllStands(w http.ResponseWriter, r *http.Request) {

	var DataS data.DataPaginate

	PageChar, _ = strconv.Atoi(r.URL.Query().Get("page"))

	CallAllStand()

	standListPaginated, paginationData := PaginateStand(StandList)
	DataS.StandList = standListPaginated
	DataS.PageData = paginationData

	temps.Temp.ExecuteTemplate(w, "AllStands", &DataS)
}

func PaginateStand(listToPaginate []data.Stand) ([]data.Stand, data.Page) {

	var PaginationStand data.Page

	perPage := 10                                 //nombre d'elements par page
	offSet := 10                                  //decalage pour l'affichage grace aux index
	nbPage := (len(listToPaginate) / perPage)     //nombre de page
	nbLastPage := (len(listToPaginate) % perPage) //nombre d'éléments dans la dernière page (doit être inférieur a nbPage)

	//assignation des numero de page par rapport a la currentPage
	PaginationStand.PreviousPage = PageChar - 1
	PaginationStand.CurrentPage = PageChar
	PaginationStand.NextPage = PageChar + 1
	PaginationStand.MaxPage = nbPage

	//changement du offset pour la derniere page lorsqu'il y a moins de 10 elements sur la dernière page.
	if PaginationStand.CurrentPage == PaginationStand.MaxPage {
		offSet = nbLastPage
	}

	//stockage de la nouvelle liste de 10 elements dans displaychar
	DisplayStand := listToPaginate[(PaginationStand.CurrentPage * perPage):(PaginationStand.CurrentPage*perPage + offSet)]

	//return de la liste de chararcters et de la struct de pagination
	return DisplayStand, PaginationStand
}

// appel et marshal le json de tous les stands et renvoi un tableau de struct
func CallAllStand() {
	urlApi := "https://stand-by-me.herokuapp.com/api/v1/stands/"

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

	json.Unmarshal(body, &StandList)
}
