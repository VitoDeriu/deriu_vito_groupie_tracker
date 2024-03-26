package controller

import (
	"GroupieTrackerJJBA/temps"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Filtre(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erreur lors de la récupération des données du formulaire", http.StatusInternalServerError)
			return
		}

		//recup des parametres=valeur choisis dans le form de filtre
		//on stock ces combinaisons p=v dans une slice de string

		//si la section de radio n'a pas été coché le form renvoi une string vide
		//qu'on ne rajoute pas a la slice
		var query []string

		if r.FormValue("chapter") != "" {
			query = append(query, "chapter="+r.FormValue("chapter"))
			fmt.Println(query)
		}
		if r.FormValue("nationality") != "" {
			query = append(query, "nationality="+r.FormValue("nationality"))
			fmt.Println(query)
		}
		if r.FormValue("human") != "" {
			query = append(query, "isHuman="+r.FormValue("human"))
			fmt.Println(query)
		}
		if r.FormValue("alive") != "" {
			query = append(query, "living="+r.FormValue("alive"))
			fmt.Println(query)
		}

		//création de l'url de call avec les p=v récup plus haut
		var urlfiltre string
		for i, c := range query {
			if i == len(query)-1 {
				urlfiltre += c
			} else {
				urlfiltre += c + "&"
			}
		}
		//appel du json filtré grace a l'url construite
		CallByFiltre(urlfiltre)
	}
	temps.Temp.ExecuteTemplate(w, "AllCharacters", ResultSearchCaracter)
}

// prend en parametre une string de combinaison de parametre=valeur pour faire un call avec ces query
// aller voir la doc de l'api pour voir quelles combinaisons sont possibles.
// ne peut pas prendre plus d'une valeur par parametre ni deux fois le meme parametre avec valeur différente.
// renvoi un tableau de Character.
func CallByFiltre(p string) {

	urlApi := "https://stand-by-me.herokuapp.com/api/v1/characters/query/query?" + p
	println("l'url de call : ", urlApi)
	println("le p : ", p)
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
