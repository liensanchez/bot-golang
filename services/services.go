package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Frase struct {
	Text string `json:"frase"`
}

func GetFrase() []Frase {

	res, err := http.Get("http://localhost:3010/api/frases")
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

	res, err := http.Get("http://localhost:3010/api/chistes")
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

	res, err := http.Get("http://localhost:3010/api/refranes")
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
