package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"struct/list/bins"
	"struct/list/config"
	"struct/list/storage"
)

func ReadKey(config config.Config) {

}

func CreateBin(file string, name string) bins.BinList {
	binList := bins.BinList{}
	data, err := storage.ReadFile(file)
	url := "https://api.jsonbin.io/v3/b"
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		return binList
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return binList
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return binList
	}
	return binList
}

func UpdateBin(file string, id string) {

}

func DeleteBin(id string) {
	url := "https://api.jsonbin.io/v3/b/66602054ad19ca34f8747e9c"
}

func GetBin(id string) {
	url := "https://api.jsonbin.io/v3/b/66602054ad19ca34f8747e9c"
}

func ListBin() {

}
