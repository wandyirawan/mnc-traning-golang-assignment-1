package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Biodata struct {
	Nama      string `json:"nama,omitempty"`
	Alamat    string `json:"alamat,omitempty"`
	Pekerjaan string `json:"pekerjaan,omitempty"`
	Alasan    string `json:"alasan,omitempty"`
}

type Biodatas struct {
	Biodata []Biodata `json:"biodata,omitempty"`
}

func main() {
	content, err := ioutil.ReadFile("./biodata.json")
	if err != nil {
		log.Fatal("error when opening file")
	}
	var payloads Biodatas

	err = json.Unmarshal(content, &payloads)
	if err != nil {
		log.Fatal("error during unmashal()")
	}
	input := os.Args[1]
	i, _ := strconv.Atoi(input)
	if i <= 0 || i > len(payloads.Biodata) {
		log.Fatal("data not found")
	}
	result, err := json.Marshal(&payloads.Biodata[i-1])
	if err != nil {
		log.Fatal("error during mashal()")
	}
	fmt.Printf("%s", string(result))
}
