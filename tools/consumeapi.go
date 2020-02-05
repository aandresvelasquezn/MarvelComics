package tools

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Comic Estructura
type Comic struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

// GetComics Get all comics list
func GetComics() {
	var urlAPI = "https://gateway.marvel.com:443/v1/public/comics?apikey=" + GetDotEnvVariable("PUBLIC_KEY") + "&ts=1&hash="

	var cliente = &http.Client{Timeout: 10 * time.Second}
	hashCode := GenerateHash()
	response, err := cliente.Get(fmt.Sprintf(urlAPI + "" + hashCode))
	fmt.Println(urlAPI + "" + hashCode)

	if err != nil {
		panic(err.Error())
	}

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

	// var comics []Comic

	// err = json.NewDecoder(data).Decode(&comics)

	// if err != nil {
	// 	panic(err.Error())
	// }go get -u github.com/golang/dep/cmd/dep

	// fmt.Println(comics)∫
}
