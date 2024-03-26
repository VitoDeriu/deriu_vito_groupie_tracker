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

var StandList []data.Stand

func AllStands(w http.ResponseWriter, r *http.Request) {

	CallAllStand()

	temps.Temp.ExecuteTemplate(w, "AllStands", StandList)
}

//appel et marshal le json de tous les stands et renvoi un tableau de struct
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
