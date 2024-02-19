package utils

func ValidateStok(stok uint, qty uint) bool {

	if stok >= qty {
		return false
	}

	return true
}
