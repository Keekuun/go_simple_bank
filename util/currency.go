package util

// 支持的货币代码
const (
	USD = "USD"
	CNY = "CNY"
	EUR = "EUR"
	CAD = "CAD"
)

// IsSupportCurrency 判断传入的货币是否支持
func IsSupportCurrency(currency string) bool {
	switch currency {
	case USD, CNY, EUR, CAD:
		return true
	}
	return false
}
