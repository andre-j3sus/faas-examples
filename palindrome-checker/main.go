package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"
)

type PalindromeRequest struct {
	Text string `json:"text"`
}

type PalindromeResponse struct {
	Text         string `json:"text"`
	IsPalindrome bool   `json:"isPalindrome"`
	Explanation  string `json:"explanation,omitempty"`
}

func isPalindrome(s string) bool {
	// Convert to lowercase and remove non-alphanumeric characters
	var cleaned []rune
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}

	// Check if palindrome
	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		if cleaned[i] != cleaned[j] {
			return false
		}
	}
	return len(cleaned) > 0
}

func handlePalindromeCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req PalindromeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		http.Error(w, "Text is required", http.StatusBadRequest)
		return
	}

	isPal := isPalindrome(req.Text)
	
	var explanation string
	if isPal {
		explanation = "The text reads the same backward as forward (ignoring case, spaces, and non-alphanumeric characters)."
	} else {
		explanation = "The text does not read the same backward as forward."
	}

	response := PalindromeResponse{
		Text:         req.Text,
		IsPalindrome: isPal,
		Explanation:  explanation,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/check", handlePalindromeCheck)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Palindrome Checker function listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}