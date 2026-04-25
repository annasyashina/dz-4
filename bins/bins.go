package bins

import (
	"encoding/json"
	"fmt"

	//"struct/list/bins"

	"struct/list/file"

	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type Account struct {
	Login     string    `json:"login"` //`json:"login" xml:"test"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BinList = []Bin

func WriteBinList(binList BinList) {

	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println("Ошибка преобразования")
		return
	}
	//fmt.Println(string(data))
	//data, err := binList

	file.WriteFile(data, "storage.json")
}

func ReadBinList(binList BinList) BinList {
	file, err := file.ReadFile("storage.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return binList
	}

	err = json.Unmarshal(file, &binList)
	if err != nil {
		fmt.Println("Ошибка разбора json файла")
		return binList
	}
	return binList
}
