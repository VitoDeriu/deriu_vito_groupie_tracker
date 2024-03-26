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

var PertinentCharacter data.Character

func DisplayCharacter(w http.ResponseWriter, r *http.Request) {

	id := (r.URL.Query().Get("id"))
	//ligne en commentaire, j'ai pas de gestion d'erreur, faut y reflechir
	// if err != nil {
	// 	http.Error(w, "ID invalide", http.StatusBadRequest)
	// 	return
	// }
	data := CallCharById(id)
	fmt.Println("la var data avec l'envoi dans le template: ", data)

	temps.Temp.ExecuteTemplate(w, "Character", data)
}

//appel seulement le Character correspondant a l'ID passé en parametre.
//renvoi une struct correspondnant au bon Characater
func CallCharById(id string) data.Character {
	urlApi := "https://stand-by-me.herokuapp.com/api/v1/characters/" + id
	println("l'url de call : ", urlApi)
	println("l'id : ", id)
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

	json.Unmarshal(body, &PertinentCharacter)
	fmt.Println("call done")
	fmt.Println(PertinentCharacter)
	return PertinentCharacter
}
