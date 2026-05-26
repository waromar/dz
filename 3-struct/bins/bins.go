package bins

import (
	"errors"
	"fmt"
	"time"
)


type Bin struct {
	ID        string `json:"id"`
	Private   bool `json:"private"`
	CreatedAt time.Time `json:"time"`
	Name      string `json:"name"`
}

type BinList struct {
	Bins []Bin
}

func InitBin(id string, private bool, name string) (*Bin, error) {
	if id == "" {
		return &Bin{}, errors.New("Ошибка ID")
	}
	if name == "" {
		return &Bin{}, errors.New("Ошибка name")
	}
	switch private {
case true:
		fmt.Println("Bin локальный")
	case false:
		fmt.Println("Bin публичный")
	}

	return &Bin{
		ID:        id,
		Private:   private,
		Name:      name,
		CreatedAt: time.Now(),
	}, nil
}

func InitBins(bins []Bin) *BinList {
	return &BinList{
		Bins: bins,
	}
}
