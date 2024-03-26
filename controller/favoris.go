package controller

import (
	"GroupieTrackerJJBA/data"
	"GroupieTrackerJJBA/temps"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Favoris(w http.ResponseWriter, r *http.Request) {

	var DisplayFav data.ResultSearch

	DisplayFavChar, err := ReadJSONChar()
	if err != nil {
		fmt.Println("Error reading Json char : ", err.Error())
		//voir pour mettre un template erreur de lecture de json
	}
	DisplayFavStand, err := ReadJSONStand()
	if err != nil {
		fmt.Println("Error reading Json stand : ", err.Error())
		//voir pour mettre un template erreur de lecture de json
	}

	DisplayFav.ResultCharacter = DisplayFavChar
	DisplayFav.ResultStand = DisplayFavStand

	data := DisplayFav

	temps.Temp.ExecuteTemplate(w, "Favoris", data)
}

func FavoriTreatment(w http.ResponseWriter, r *http.Request) {

	id := (r.URL.Query().Get("id"))
	tipe := (r.URL.Query().Get("type"))

	if tipe == "manieur" {
		ListFavChar, err := ReadJSONChar()
		if err != nil {
			fmt.Println("Error reading json char : ", err.Error())
			//voir pour mettre un template erreur de lecture de json
		}

		for _, c := range ListFavChar {
			if c.Id == id {
				temps.Temp.ExecuteTemplate(w, "AllCharacters", CharacterList)
				return
			}
		}
		ListFavChar = append(ListFavChar, CallCharById(id))
		EditJSONChar(ListFavChar)
		temps.Temp.ExecuteTemplate(w, "AllCharacters", CharacterList)
	} else {
		ListFavStand, err := ReadJSONStand()
		if err != nil {
			fmt.Println("Error reading json stand : ", err.Error())
			//voir pour mettre un template erreur de lecture de json
		}

		for _, c := range ListFavStand {
			if c.Id == id {
				temps.Temp.ExecuteTemplate(w, "AllStands", StandList)
				return
			}
		}
		ListFavStand = append(ListFavStand, CallStandById(id))
		EditJSONStand(ListFavStand)
		temps.Temp.ExecuteTemplate(w, "AllStands", StandList)
	}

}

func FavoriDelete(w http.ResponseWriter, r *http.Request) {

	id := (r.URL.Query().Get("id"))
	tipe := (r.URL.Query().Get("type"))

	if tipe == "character"{

		ListFavCharToDel, err := ReadJSONChar()
		if err != nil {
			fmt.Println("Error reading json char : ", err.Error())
			//voir pour mettre un template erreur de lecture de json
		}

		var ListNewFav []data.Character

		for _, c := range ListFavCharToDel {
			if c.Id != id {
				ListNewFav = append(ListNewFav,c)
			}
		}

		// for i, _ := range ListFavCharToDel{
		// 	if ListFavCharToDel[i].Id == id{
		// 		ListFavCharToDel = append(ListFavCharToDel[:i], ListFavCharToDel[i+1:]... )
		// 	}
		// }

		EditJSONChar(ListNewFav)

	} else {

		ListFavStandToDel, err := ReadJSONStand()
		if err != nil {
			fmt.Println("Error reading json stand : ", err.Error())
			//voir pour mettre un template erreur de lecture de json
		}

		var ListNewFavStand []data.Stand

		for _, d := range ListFavStandToDel {
			if d.Id != id {
				ListNewFavStand = append(ListNewFavStand, d)
			}
		}

		// for i, _ := range ListFavCharToDel{
		// 	if ListFavCharToDel[i].Id == id{
		// 		ListFavCharToDel = append(ListFavCharToDel[:i], ListFavCharToDel[i+1:]... )
		// 	}
		// }

		EditJSONStand(ListNewFavStand)
	}
	
	http.Redirect(w, r, "/favoris", http.StatusMovedPermanently)
}


// Fonction pour mettre le JSON dans une struct
func ReadJSONChar() ([]data.Character, error) {
	fmt.Println("JSON en cours de lecture...")
	jsonFile, err := os.ReadFile("data/favChar.json")
	if err != nil {
		fmt.Println("Error reading json char : ", err.Error())
	}

	var jsonData []data.Character
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err

}

func ReadJSONStand() ([]data.Stand, error) {
	fmt.Println("JSON en cours de lecture...")
	jsonFile, err := os.ReadFile("data/favStand.json")
	if err != nil {
		fmt.Println("Error reading json stand : ", err.Error())
	}

	var jsonData []data.Stand
	err = json.Unmarshal(jsonFile, &jsonData)
	return jsonData, err
}

// Fonction pour modifier le JSON
func EditJSONChar(ModifiedFav []data.Character) {
	fmt.Println("JSON en cours de modification...")
	modifiedJSON, errMarshal := json.Marshal(ModifiedFav)
	if errMarshal != nil {
		fmt.Println("Error encodage json char : ", errMarshal.Error())
		return
	}

	// Écrire le JSON modifié dans le fichier
	if err := os.WriteFile("data/favChar.json", modifiedJSON, 0644); err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié: ", err)
	} else {
		fmt.Println("JSON modified successfully")
	}
}

func EditJSONStand(ModifiedFav []data.Stand) {
	fmt.Println("JSON en cours de modification...")
	modifiedJSON, errMarshal := json.Marshal(ModifiedFav)
	if errMarshal != nil {
		fmt.Println("Error encodage json stand : ", errMarshal.Error())
		return
	}

	// Écrire le JSON modifié dans le fichier
	if err := os.WriteFile("data/favStand.json", modifiedJSON, 0644); err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier JSON modifié: ", err)
	} else {
		fmt.Println("JSON modified successfully")
	}
}
