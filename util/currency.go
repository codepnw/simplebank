package util

const (
	THB = "THB"
	USD = "USD"
)

func IsSupportCurrency(currency string) bool {
	switch currency {
	case THB, USD:
		return true
	}
	return false
}