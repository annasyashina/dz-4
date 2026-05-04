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

type BinResponse struct {
	Metadata struct {
		ID string `json:"id"`
	} `json:"metadata"`
	Record interface{} `json:"record"`
}

func ReadKey(config config.Config) string {
	return config.Key
}

func CreateBin(file string, name string) BinResponse {
	binList := BinResponse{}
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
	client := &http.Client{}
	response, err := client.Do(resp)
	if err != nil {
		return binList
	}
	defer response.Body.Close()
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return binList
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return binList
	}
	return binList
}

func UpdateBin(file string, id string) BinResponse {
	binList := BinResponse{}
	data, err := storage.ReadFile(file)

	url := fmt.Sprintf("%s/b/%s", "https://api.jsonbin.io/v3/", id)
	resp, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		return binList
	}
	cfg := &config.Config{Key: os.Getenv("KEY")}

	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("X-Master-Key", cfg.GetAPIKey())
	client := &http.Client{}
	response, err := client.Do(resp)
	if err != nil {
		return binList
	}
	defer response.Body.Close()
	data, err = io.ReadAll(response.Body)
	if err != nil {
		return binList
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return binList
	}
	return binList
}

func DeleteBin(id string) error {
	//url := "https://api.jsonbin.io/v3/b/66602054ad19ca34f8747e9c"
	binList := bins.BinList{}
	url := fmt.Sprintf("%s/b/%s", "https://api.jsonbin.io/v3/", id)
	resp, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	cfg := &config.Config{Key: os.Getenv("KEY")}

	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("X-Master-Key", cfg.GetAPIKey())
	client := &http.Client{}
	response, err := client.Do(resp)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return err
	}
	return err
}

func GetBin(id string) map[string]interface{} {

	url := fmt.Sprintf("%s/b/%s", "https://api.jsonbin.io/v3/", id)
	resp, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil
	}
	cfg := &config.Config{Key: os.Getenv("KEY")}

	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("X-Master-Key", cfg.GetAPIKey())
	client := &http.Client{}
	response, err := client.Do(resp)
	if err != nil {
		return nil
	}
	defer response.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil
	}

	return result
}

func ListBin() {

}
