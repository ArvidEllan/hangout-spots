package services

import (
	"fmt"
	"time"
)

// InitiateSTK simulates an STK push; replace with real Daraja integration.
func InitiateSTK(phone string, amount int, reference string) (string, error) {
	if phone == "" || amount <= 0 {
		return "", fmt.Errorf("invalid stk request")
	}
	return fmt.Sprintf("MOCK-%d", time.Now().UnixNano()), nil
}

