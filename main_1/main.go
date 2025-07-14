package main

import (
	"fmt"
	"go-core-task/main_1/handlers"
)

func main() {
	variables := handlers.CreateVariables()

	handlers.PrintTypeInfo("decimal", variables.Decimal)
	handlers.PrintTypeInfo("octal", variables.Octal)
	handlers.PrintTypeInfo("hexadecimal", variables.Hexadecimal)
	handlers.PrintTypeInfo("floatNum", variables.FloatNum)
	handlers.PrintTypeInfo("message", variables.Message)
	handlers.PrintTypeInfo("isTrue", variables.IsTrue)
	handlers.PrintTypeInfo("complexNum", variables.ComplexNum)

	stringRepresentations := handlers.ConvertToStrings(&variables)

	result := handlers.BuildResultString(stringRepresentations)

	torunes := handlers.ConvertToRune(result)

	hash := handlers.HashRunesDirectly(torunes)

	fmt.Println("\nрезультат хэша:")
	fmt.Println(hash)
	fmt.Println("\nОбъединенная строка:")
	fmt.Print(result)
}
