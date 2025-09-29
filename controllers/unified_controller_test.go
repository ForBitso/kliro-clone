package controllers

import (
	"context"
	"encoding/base64"
	"testing"
)

func TestGeneratePaymePayment_GET_Fallback(t *testing.T) {
	// Test parameters based on requirements
	merchantID := "646c8bff2cb83937a7551c95"
	anketaID := int64(778957)
	amountUZS := int64(117600)
	returnURL := "https://ersp.e-osgo.uz/uz/site/export-to-pdf?id=778957"
	orderKeyName := "ac.key"

	// Expected decoded string from requirements
	expectedDecoded := "https://ersp.e-osgo.uz/uz/site/export-to-pdf?id=778957;ct=20000;m=646c8bff2cb83937a7551c95;ac.key=778957;a=11760000"

	// Generate Payme URL using GET fallback (usePost = false)
	paymeURL, rawParams, err := GeneratePaymePayment(
		context.Background(),
		merchantID,
		"", // merchantKey - empty for GET
		anketaID,
		amountUZS,
		returnURL,
		orderKeyName,
		false, // usePost = false
		"",    // postURL = empty
	)

	// Check for errors
	if err != nil {
		t.Fatalf("GeneratePaymePayment returned error: %v", err)
	}

	// Check that URL is not empty
	if paymeURL == "" {
		t.Fatal("Generated Payme URL is empty")
	}

	// Check that rawParams matches expected decoded string
	actualDecoded := string(rawParams)
	if actualDecoded != expectedDecoded {
		t.Errorf("Decoded params mismatch:\nExpected: %s\nActual:   %s", expectedDecoded, actualDecoded)
	}

	// Check that URL has correct prefix
	expectedPrefix := "https://checkout.paycom.uz/"
	if len(paymeURL) <= len(expectedPrefix) || paymeURL[:len(expectedPrefix)] != expectedPrefix {
		t.Errorf("Payme URL doesn't have correct prefix. Got: %s", paymeURL)
	}

	// Extract and decode base64 part
	base64Part := paymeURL[len(expectedPrefix):]
	decoded, err := base64.StdEncoding.DecodeString(base64Part)
	if err != nil {
		t.Fatalf("Failed to decode base64 part: %v", err)
	}

	// Verify decoded content matches expected
	if string(decoded) != expectedDecoded {
		t.Errorf("Base64 decoded content mismatch:\nExpected: %s\nActual:   %s", expectedDecoded, string(decoded))
	}

	t.Logf("✅ Test passed!")
	t.Logf("Generated URL: %s", paymeURL)
	t.Logf("Decoded params: %s", actualDecoded)
}

func TestGeneratePaymePayment_OrderID_Variant(t *testing.T) {
	// Test with ac.order_id instead of ac.key
	merchantID := "646c8bff2cb83937a7551c95"
	anketaID := int64(778957)
	amountUZS := int64(117600)
	returnURL := "https://ersp.e-osgo.uz/uz/site/export-to-pdf?id=778957"
	orderKeyName := "ac.order_id"

	// Expected decoded string with ac.order_id
	expectedDecoded := "https://ersp.e-osgo.uz/uz/site/export-to-pdf?id=778957;ct=20000;m=646c8bff2cb83937a7551c95;ac.order_id=778957;a=11760000"

	// Generate Payme URL
	paymeURL, rawParams, err := GeneratePaymePayment(
		context.Background(),
		merchantID,
		"",
		anketaID,
		amountUZS,
		returnURL,
		orderKeyName,
		false,
		"",
	)

	if err != nil {
		t.Fatalf("GeneratePaymePayment returned error: %v", err)
	}

	actualDecoded := string(rawParams)
	if actualDecoded != expectedDecoded {
		t.Errorf("Decoded params mismatch:\nExpected: %s\nActual:   %s", expectedDecoded, actualDecoded)
	}

	t.Logf("✅ Test with ac.order_id passed!")
	t.Logf("Generated URL: %s", paymeURL)
	t.Logf("Decoded params: %s", actualDecoded)
}

func TestGeneratePaymePayment_InvalidInput(t *testing.T) {
	tests := []struct {
		name       string
		merchantID string
		anketaID   int64
		amountUZS  int64
		expectErr  bool
	}{
		{
			name:       "Empty merchant ID",
			merchantID: "",
			anketaID:   123,
			amountUZS:  1000,
			expectErr:  true,
		},
		{
			name:       "Zero anketa ID",
			merchantID: "test",
			anketaID:   0,
			amountUZS:  1000,
			expectErr:  true,
		},
		{
			name:       "Zero amount",
			merchantID: "test",
			anketaID:   123,
			amountUZS:  0,
			expectErr:  true,
		},
		{
			name:       "Negative amount",
			merchantID: "test",
			anketaID:   123,
			amountUZS:  -100,
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := GeneratePaymePayment(
				context.Background(),
				tt.merchantID,
				"",
				tt.anketaID,
				tt.amountUZS,
				"https://example.com",
				"ac.key",
				false,
				"",
			)

			if tt.expectErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
		})
	}
}
