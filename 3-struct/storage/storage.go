package storage

import (
	"demo/3-struct/bins"
	"encoding/json"
	"fmt"
	"os"
)

// Сохранение bin в виде json в локальном файле

func SaveBin(b bins.Bin) {
	bl := ReadBins()
	bl.Bins = append(bl.Bins, b)

	data, err := json.MarshalIndent(bl, "", "  ")
	if err != nil {
		fmt.Println( err)
		return
	}

	err = os.WriteFile("bins.json", data, 0644)
	if err != nil {
		fmt.Println( err)
	}
}
// Чтение списка bin в виде json из локального файла
func ReadBins() bins.BinList {
	var huylist bins.BinList

	_, err := os.Stat("bins.json")
	if os.IsNotExist(err) {
		return huylist
	}

	data, err := os.ReadFile("bins.json")
	if err != nil {
		fmt.Println(err)
		return huylist
	}

	err = json.Unmarshal(data, &huylist)
	if err != nil {
		fmt.Println(err)
		return huylist
	}

	return huylist
}