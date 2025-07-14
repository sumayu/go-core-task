package main

import (
	"go-core-task/main_1/handlers"
	"go-core-task/main_1/model"
	"reflect"
	"testing"
)

func TestCreateVariables(t *testing.T) {
	expected := model.Variables{
		Decimal:     42,
		Octal:       052,
		Hexadecimal: 0x2A,
		FloatNum:    3.14,
		Message:     "Golang",
		IsTrue:      true,
		ComplexNum:  complex64(1 + 2i),
	}

	got := handlers.CreateVariables()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("CreateVariables() = %v, want %v", got, expected)
	}
}

func TestConvertToStrings(t *testing.T) {
	v := model.Variables{
		Decimal:     42,
		Octal:       052,
		Hexadecimal: 0x2A,
		FloatNum:    3.14,
		Message:     "Golang",
		IsTrue:      true,
		ComplexNum:  complex64(1 + 2i),
	}

	got := handlers.ConvertToStrings(&v)

	expected := map[string]string{
		"decimal":     "42",
		"octal":       "52",
		"hexadecimal": "2a",
		"float":       "3.14",
		"string":      "Golang",
		"bool":        "true",
		"complex":     "(1.0+2.0i)",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ConvertToStrings() = %v, want %v", got, expected)
	}
}

func TestBuildResultString(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		expected string
	}{
		{
			"empty map",
			map[string]string{},
			"",
		},
		{
			"real data",
			map[string]string{
				"decimal":     "42",
				"octal":       "52",
				"hexadecimal": "2a",
			},
			"42522a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handlers.BuildResultString(tt.input); got != tt.expected {
				t.Errorf("BuildResultString() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestConvertToRune(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []rune
	}{
		{
			"empty string",
			"",
			[]rune{},
		},
		{
			"ASCII chars",
			"abc",
			[]rune{'a', 'b', 'c'},
		},
		{
			"Unicode chars",
			"Привет",
			[]rune{'П', 'р', 'и', 'в', 'е', 'т'},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handlers.ConvertToRune(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ConvertToRune() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestHashRunesDirectly(t *testing.T) {
	tests := []struct {
		name     string
		input    []rune
		expected string
	}{
		{
			"empty runes",
			[]rune{},
			"5f70bf",
		},
		{
			"simple runes",
			[]rune{'a', 'b', 'c'},
			"a3f4c6",
		},
		{
			"unicode runes",
			[]rune{'П', 'р', 'и'},
			"c3b1b4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := handlers.HashRunesDirectly(tt.input)
			if len(got) != 64 {
				t.Errorf("HashRunesDirectly() length = %d, want 64", len(got))
			}
		})
	}
}
