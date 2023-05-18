package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Frase struct {
	Text string `json:"frase"`
}

func GetFrase() []Frase {
	URL := os.Getenv("URL")

	res, err := http.Get(URL + "/frases")
	if err != nil {
		fmt.Println("error")
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("error")
	}

	var frases []Frase
	err = json.Unmarshal(body, &frases)
	if err != nil {
		fmt.Println("error")
	}

	return frases
}

func GetChiste() []Frase {
	URL := os.Getenv("URL")

	res, err := http.Get(URL + "/chistes")
	if err != nil {
		fmt.Println("error")
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("error")
	}

	var frases []Frase
	err = json.Unmarshal(body, &frases)
	if err != nil {
		fmt.Println("error")
	}

	return frases
}

func GetRefran() []Frase {
	URL := os.Getenv("URL")

	res, err := http.Get(URL + "/refranes")
	if err != nil {
		fmt.Println("error")
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println("error")
	}

	var frases []Frase
	err = json.Unmarshal(body, &frases)
	if err != nil {
		fmt.Println("error")
	}

	return frases
}
