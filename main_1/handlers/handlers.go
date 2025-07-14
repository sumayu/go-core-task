package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"go-core-task/main_1/model"
	"strconv"
	"strings"
)

func ConvertToStrings(v *model.Variables) map[string]string {
	return map[string]string{
		"decimal":     strconv.Itoa(v.Decimal),
		"octal":       strconv.FormatInt(int64(v.Octal), 8),
		"hexadecimal": strconv.FormatInt(int64(v.Hexadecimal), 16),
		"float":       strconv.FormatFloat(v.FloatNum, 'f', 2, 64),
		"string":      v.Message,
		"bool":        strconv.FormatBool(v.IsTrue),
		"complex":     fmt.Sprintf("%.1f", complex64(v.ComplexNum)),
	}
}

func ConvertToRune(r string) []rune {
	return []rune(r)
}

func HashRunesDirectly(runes []rune) string {
	salt := "go-2024"

	str := string(runes)

	mid := len(str) / 2
	saltedStr := str[:mid] + salt + str[mid:]

	hasher := sha256.New()
	hasher.Write([]byte(saltedStr))
	return hex.EncodeToString(hasher.Sum(nil))
}

func BuildResultString(data map[string]string) string {
	var builder strings.Builder
	builder.Grow(200)

	for _, v := range data {
		builder.WriteString(v)
	}

	return builder.String()
}

func PrintTypeInfo(name string, value interface{}) {
	fmt.Printf("%s (%v) имеет тип: %T\n", name, value, value)
}

func CreateVariables() model.Variables {
	return model.Variables{
		Decimal:     42,
		Octal:       052,
		Hexadecimal: 0x2A,
		FloatNum:    3.14,
		Message:     "Golang",
		IsTrue:      true,
		ComplexNum:  complex64(1 + 2i),
	}
}
