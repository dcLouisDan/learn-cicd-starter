package auth


import (
  "net/http"
  "testing"
)

func TestGetApiKey(t *testing.T) {
  headers := http.Header{}
	headers.Set("Authorization", "ApiKey my-secret-key is wrong")

	key, err := GetAPIKey(headers)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if key != "my-secret-key" {
		t.Fatalf("expected 'my-secret-key', got %q", key)
	}
}


