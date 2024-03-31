	package data

type ResultSearch struct {
	ResultCharacter []Character
	ResultStand     []Stand
}

//struct pour l'envoi dans les templates
type DataPaginate struct {
	CharList []Character
	StandList []Stand
	PageData  Page
}

// struct pour le systeme de pagination
type Page struct {
	CurrentPage  int
	NextPage     int
	PreviousPage int
	MaxPage      int
}

type Stand struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	AlternateName string `json:"alternateName"`
	JapaneseName  string `json:"japaneseName"`
	Image         string `json:"image"`
	StandUser     string `json:"standUser"`
	Chapter       string `json:"chapter"`
	Abilities     string `json:"abilities"`
	Battlecry     string `json:"battlecry"`
}

type Character struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	JapaneseName string `json:"japaneseName"`
	Image        string `json:"image"`
	Abilities    string `json:"abilities"`
	Nationality  string `json:"nationality"`
	Catchphrase  string `json:"catchphrase"`
	Family       string `json:"family"`
	Chapter      string `json:"chapter"`
	Living       bool   `json:"living"`
	IsHuman      bool   `json:"isHuman"`
}