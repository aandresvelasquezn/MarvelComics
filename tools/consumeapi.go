package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	comic "github.com/aandresvelasquezn/MarvelComics/internal/Model/Comic"
)

// Comics Get all comics list
func Comics() {
	var urlAPI = DotEnvVariable("URL_REQUEST_COMIC") + "?apikey=" + DotEnvVariable("PUBLIC_KEY") + "&ts=1&hash="

	var cliente = &http.Client{Timeout: 10 * time.Second}
	hashCode := GenerateHash()
	response, err := cliente.Get(fmt.Sprintf(urlAPI + "" + hashCode))
	fmt.Println(urlAPI + "" + hashCode)

	if err != nil {
		panic(err.Error())
	}

	var comics comic.DataWrapper

	err = json.NewDecoder(response.Body).Decode(&comics)

	if err != nil {
		panic(err.Error())
	}
	fmt.Println(comics)
}

//ComicByID Get comic data by Id
func ComicByID(comicID int) {

	comID := strconv.Itoa(comicID)
	var urlAPI = DotEnvVariable("URL_REQUEST_COMIC") + "/" + comID + "?apikey=" + DotEnvVariable("PUBLIC_KEY") + "&ts=1&hash="

	var cliente = &http.Client{Timeout: 10 * time.Second}
	hashCode := GenerateHash()
	response, err := cliente.Get(fmt.Sprintf(urlAPI + "" + hashCode))
	fmt.Println(urlAPI + "" + hashCode)

	if err != nil {
		panic(err.Error())
	}

	var comic comic.DataWrapper

	err = json.NewDecoder(response.Body).Decode(&comic)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(comic)
}
