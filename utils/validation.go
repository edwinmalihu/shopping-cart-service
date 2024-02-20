package utils

type StokData struct {
	Stok  uint `json:"stok"`
	State bool `json:"state"`
}

func ValidateStok(stok uint, qty uint) StokData {
	var res StokData
	if stok >= qty {
		res = StokData{
			Stok:  0,
			State: false,
		}

		return res
	}

	nilaiStok := stok - qty

	res = StokData{
		Stok:  nilaiStok,
		State: true,
	}

	return res
}

func TotalPrice(qty int, price float64) float64 {
	total := qty * int(price)

	return float64(total)
}
