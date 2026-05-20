package struct


type Bin struct {
	ID        string
	Private   bool
	CreatedAt time.Time
	Name      string
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
	if private == true {
		fmt.Println("Bin локальный")
	} else if private == false {
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
