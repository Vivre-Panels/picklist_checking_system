package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"picklist_checking_system/service"

	"github.com/joho/godotenv"
)

func GenrateAuthUrl(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env")
	}

	ZohoAuthUrl := os.Getenv("ZOHO_OAUTH_GENRATE_URL")
	clientID := os.Getenv("ZOHO_CLIENT_ID")
	RedirectUrl := os.Getenv("ZOHO_REDIRECT_URI")
	ZohoScope := os.Getenv("ZOHO_SCOPE")

	if ZohoAuthUrl == "" || clientID == "" || RedirectUrl == "" || ZohoScope == "" {
		http.Error(w, "Missing Zoho OAuth config", http.StatusInternalServerError)
		return
	}

	query := url.Values{}
	query.Set("response_type", "code")
	query.Set("client_id", clientID)
	query.Set("scope", ZohoScope)
	query.Set("redirect_uri", RedirectUrl)
	query.Set("access_type", "offline")

	fullURL := fmt.Sprintf("%s?%s", ZohoAuthUrl, query.Encode())
	w.Write([]byte(fullURL))
}

func GenrateAuthToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request Method", http.StatusMethodNotAllowed)
		return
	}

	oauth_code := r.URL.Query().Get("code")
	if oauth_code == "" {
		http.Error(w, "Missing code query parameter", http.StatusBadRequest)
		return
	}

	expiery := service.GenrateToken(oauth_code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":    true,
		"message":    "OAuth authorization completed successfully!",
		"expires_in": expiery,
	})
}
