package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"struct/list/bins"
	"struct/list/config"
	"struct/list/storage"
)

func ReadKey(config config.Config) string {
	return config.Key
}

func CreateBin(file string, name string) bins.BinList {
	binList := bins.BinList{}
	data, err := storage.ReadFile(file)
	url := "https://api.jsonbin.io/v3/b"

	//resp, err := http.Post(url, headers, bytes.NewBuffer(data))

	resp, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return binList
	}

	cfg := &config.Config{Key: os.Getenv("KEY")}

	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("X-Master-Key", cfg.GetAPIKey())
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

func UpdateBin(file string, id string) bins.BinList {
	binList := bins.BinList{}
	data, err := storage.ReadFile(file)

	url := fmt.Sprintf("%s/bins/%s", "https://api.jsonbin.io/v3/", id)
	resp, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return binList
	}
	cfg := &config.Config{Key: os.Getenv("KEY")}

	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("X-Master-Key", cfg.GetAPIKey())
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

func DeleteBin(id string) {
	//url := "https://api.jsonbin.io/v3/b/66602054ad19ca34f8747e9c"
	binList := bins.BinList{}
	url := fmt.Sprintf("%s/bins/%s", "https://api.jsonbin.io/v3/", id)
	resp, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}
	cfg := &config.Config{Key: os.Getenv("KEY")}

	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("X-Master-Key", cfg.GetAPIKey())
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return
	}
}

func GetBin(id string) map[string]interface{} {

	url := fmt.Sprintf("%s/bins/%s", "https://api.jsonbin.io/v3/", id)
	resp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	cfg := &config.Config{Key: os.Getenv("KEY")}

	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("X-Master-Key", cfg.GetAPIKey())
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil
	}

	return result
}

func ListBin() {

}
