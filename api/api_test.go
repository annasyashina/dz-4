package api_test

import (
	"encoding/json"
	"os"
	"reflect"
	"struct/list/api"
	"testing"
)

func TestCreateBin(t *testing.T) {
	file := "my.json"
	name := "my-bin"
	expected := api.BinResponse{}
	data, err := os.ReadFile(file)
	if err != nil {
		err = json.Unmarshal(data, &expected)
	}

	result := api.CreateBin(file, name)
	if result != expected {
		t.Errorf("Получено %v, ожидалось %v", result, expected)
	}
	id := result.Metadata.ID
	api.DeleteBin(id)
}

func TestUpdateBin(t *testing.T) {
	file := "my.json"
	name := "my-bin"
	expected := api.BinResponse{}
	data, err := os.ReadFile(file)
	if err != nil {
		err = json.Unmarshal(data, &expected)
	}

	result := api.CreateBin(file, name)

	id := result.Metadata.ID

	result_update := api.UpdateBin(file, id)
	if result_update != expected {
		t.Errorf("Получено %v, ожидалось %v", result, expected)
	}

	api.DeleteBin(id)
}

func TestGetBin(t *testing.T) {
	file := "my.json"
	name := "my-bin"

	var expected map[string]interface{}

	data, err := os.ReadFile(file)
	if err != nil {
		err = json.Unmarshal(data, &expected)
	}

	result := api.CreateBin(file, name)

	id := result.Metadata.ID

	result_get := api.GetBin(id)
	//if result_get != expected {
	if !reflect.DeepEqual(result_get, expected) {
		t.Errorf("Получено %v, ожидалось %v", result, expected)
	}

	api.DeleteBin(id)
}

func TestDeleteBin(t *testing.T) {
	file := "my.json"
	name := "my-bin"
	var expected error
	expected = nil

	result := api.CreateBin(file, name)

	id := result.Metadata.ID

	result_delete := api.DeleteBin(id)
	if result_delete != expected {
		t.Errorf("Получено %v, ожидалось %v", result, expected)
	}
}
