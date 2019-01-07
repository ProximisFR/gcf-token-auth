package gcftokenauth

import (
	"fmt"
	"net/http"
	"os"
)

// Auth Compare provided token with function environment variable to allow Google Cloud function to execute or not
func Auth(w http.ResponseWriter, r *http.Request) bool {
	fAPIKey := os.Getenv("X_PROXIMIS_API_KEY")
	if fAPIKey == "" {
		http.Error(w, "API key env var missing", http.StatusInternalServerError)
		return false
	}

	APIKey := r.Header.Get("X-Proximis-Api-Key")

	fmt.Print(APIKey)
	if APIKey == "" || APIKey != fAPIKey {
		http.Error(w, "Missing or wrong api key header", http.StatusUnauthorized)
		return false
	}
	return true
}
