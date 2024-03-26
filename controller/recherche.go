package controller

import (
	"GroupieTrackerJJBA/data"
	"GroupieTrackerJJBA/temps"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var ResultSearchCaracter []data.Character
var ResultSearchStand []data.Stand
var ResultSearch data.ResultSearch

func DisplayRecherche(w http.ResponseWriter, r *http.Request) {
	temps.Temp.ExecuteTemplate(w, "Recherche", nil)
}

func Recherche(w http.ResponseWriter, r *http.Request) {

	Query := r.FormValue("search")
	Query = strings.Replace(Query, " ", "%20", -1)

	if r.Method == "POST" {

		CallCharbyname(Query) //modif la variable
		CallStandbyname(Query)
		ResultSearch.ResultCharacter = ResultSearchCaracter
		ResultSearch.ResultStand = ResultSearchStand
		data := ResultSearch
		temps.Temp.ExecuteTemplate(w, "SearchResult", data)
	}
}

// Fonction de recherche par Nom

func CallCharbyname(c string) {

	urlApi := "https://stand-by-me.herokuapp.com/api/v1/characters/query/query?name=" + c
	println("l'url de call : ", urlApi)
	println("le c : ", c)
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

	json.Unmarshal(body, &ResultSearchCaracter)
	fmt.Println("call done")
	fmt.Println(ResultSearchCaracter)
}

func CallStandbyname(c string) {

	urlApi := "https://stand-by-me.herokuapp.com/api/v1/stands/query/query?name=" + c
	println("l'url de call : ", urlApi)
	println("le c : ", c)
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

	json.Unmarshal(body, &ResultSearchStand)
	fmt.Println("call done")
	fmt.Println(ResultSearchStand)
}
