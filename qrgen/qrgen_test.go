package qrgen

import (
	"testing"
)

func TestGenerateQRCodeShouldReturnValue(t *testing.T){
	result := GenerateQRCode("4545434u-3432894")

	if result == nil {
		t.Errorf("Nil is returned")
	}

	if len(result) == 0 {
		t.Errorf("Function does not return any data")
	}
}