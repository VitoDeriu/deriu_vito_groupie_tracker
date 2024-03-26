package controller

import (
	"GroupieTrackerJJBA/data"
	"GroupieTrackerJJBA/temps"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var PertinentStand data.Stand

func DisplayStand(w http.ResponseWriter, r *http.Request) {

	id := (r.URL.Query().Get("id"))

	data := CallStandById(id)

	temps.Temp.ExecuteTemplate(w, "Stand", data)
}

func CallStandById(id string) data.Stand {

	urlApi := "https://stand-by-me.herokuapp.com/api/v1/stands/" + id
	println(urlApi)
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

	json.Unmarshal(body, &PertinentStand)
	fmt.Println("call done")
	fmt.Println(PertinentStand)
	return PertinentStand
}
