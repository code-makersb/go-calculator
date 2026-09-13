package calculator

import (
	"testing"
)

// 1. Add функциясын тексеру
func TestAdd(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected int
	}{
		{name: "Оң сандарды қосу", num1: 2, num2: 3, expected: 5},
		{name: "Теріс сандарды қосу", num1: -1, num2: -5, expected: -6},
		{name: "Нөлді қосу", num1: 5, num2: 0, expected: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.num1, tt.num2)
			if result != tt.expected {
				t.Errorf("%s үшін қате: алып жатқанымыз %d, болуы тиіс %d", tt.name, result, tt.expected)
			}
		})
	}
}

// 2. Divide (Бөлу) функциясын тексеру
func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		num1        int
		num2        int
		expected    int
		expectError bool
	}{
		{name: "Дұрыс бөлу", num1: 6, num2: 2, expected: 3, expectError: false},
		{name: "Нөлге бөлу қатесі", num1: 5, num2: 0, expected: 0, expectError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.num1, tt.num2)

			if tt.expectError && err == nil {
				t.Errorf("%s үшін қате: Нөлге бөлу қатесі шығуы тиіс еді!", tt.name)
			}

			if !tt.expectError && err != nil {
				t.Errorf("%s үшін қате: күтілмеген қате шықты: %v", tt.name, err)
			}

			if result != tt.expected {
				t.Errorf("%s үшін қате: алып жатқанымыз %d, болуы тиіс %d", tt.name, result, tt.expected)
			}
		})
	}
}

// 3. Power (Дәрежеге шығару) функциясын тексеру
func TestPower(t *testing.T) {
	tests := []struct {
		name        string
		base        int
		exponent    int
		expected    int
		expectError bool
	}{
		{name: "Оң дәреже", base: 2, exponent: 3, expected: 8, expectError: false},
		{name: "Нөлдік дәреже", base: 5, exponent: 0, expected: 1, expectError: false},
		{name: "Теріс дәреже қатесі", base: 2, exponent: -1, expected: 0, expectError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Power(tt.base, tt.exponent)

			if tt.expectError && err == nil {
				t.Errorf("%s үшін қате: теріс дәреже қатесі шығуы тиіс еді!", tt.name)
			}

			if !tt.expectError && err != nil {
				t.Errorf("%s үшін қате: күтілмеген қате шықты: %v", tt.name, err)
			}

			if result != tt.expected {
				t.Errorf("%s үшін қате: алып жатқанымыз %d, болуы тиіс %d", tt.name, result, tt.expected)
			}
		})
	}
}
