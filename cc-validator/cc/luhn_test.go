package cc

import "testing"

func TestIsValidLugn(t *testing.T) {
	tests := []struct {
		name       string
		cardNumber string
		want       bool
	}{
		{"valid card number", "1234567890123452", true},
		{"invalid card number", "1234567890123456", false},
		{"invalid card number", "12345678901234567", false},
		{"invalid card number", "123456789012345", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidLuhn(tt.cardNumber); got != tt.want {
				t.Errorf("isValidLuhn() = %v, want %v", got, tt.want)
			}
		})
	}
}
