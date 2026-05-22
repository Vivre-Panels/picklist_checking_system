package handler

import (
	"encoding/json"
	"net/http"
	service "picklist_checking_system/service/webhook"
)

func BookWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	defer r.Body.Close()

	var booksPayload map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&booksPayload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// fmt.Println("Received:", booksPayload)
	service.HandleWebhook(booksPayload)

	// ---- Response ----
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Webhook Post Successful",
		"response": 200,
	})
	
	if err != nil {
		http.Error(w, "Response encoding failed", http.StatusInternalServerError)
	}
}
