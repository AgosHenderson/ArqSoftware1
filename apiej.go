package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Categories []Category

type Category struct {
	ID   string `json:"id"`
	Nmae string `json:"name"`
}

func main() {
	cats, err := GetCategories("MLA")
	if err != nil {
		//validar
	}
	fmt.Println("Las categorias son...")
}

func GetCategories(siteID string) (Categories, error) {

	response, err := http.Get               //completar
	bytes := ioutil.ReadAll(response.Bytes) //completar
	var cats Categories
	err := json.Unmarshal(bytes, &cats)
	return cats, nil

}
