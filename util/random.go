package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt 返回 [min, max] 之间的随机数
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString 返回 长度为 n 的随机字符串
func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner 生成随机 用户名
func RandomOwner() string {
	return RandomString(int(RandomInt(4, 8)))
}

// RandomMoney 生成随机 金额
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency 生成随机 货币代码
func RandomCurrency() string {
	currencies := []string{USD, CNY, EUR, CAD}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail 生成随机 email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
