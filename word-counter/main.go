package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type WordCountRequest struct {
	Text string `json:"text"`
}

type WordCountResponse struct {
	Text         string            `json:"text"`
	WordCount    int               `json:"wordCount"`
	CharCount    int               `json:"charCount"`
	FrequentWords map[string]int   `json:"frequentWords"`
}

func countWords(text string) (int, map[string]int) {
	// Simple word counting: split by whitespace
	if text == "" {
		return 0, make(map[string]int)
	}
	
	// Convert to lowercase for better word frequency analysis
	text = strings.ToLower(text)
	
	// Split text into words, removing punctuation
	re := regexp.MustCompile(`[^\p{L}\p{N}]+`)
	words := re.Split(text, -1)
	
	// Filter out empty strings
	var filteredWords []string
	for _, word := range words {
		if word != "" {
			filteredWords = append(filteredWords, word)
		}
	}
	
	// Count word frequency
	wordFreq := make(map[string]int)
	for _, word := range filteredWords {
		wordFreq[word]++
	}
	
	return len(filteredWords), wordFreq
}

func getTopWords(wordFreq map[string]int, limit int) map[string]int {
	if len(wordFreq) <= limit {
		return wordFreq
	}
	
	type wordFreqPair struct {
		word string
		freq int
	}
	
	// Convert map to slice for sorting
	pairs := make([]wordFreqPair, 0, len(wordFreq))
	for word, freq := range wordFreq {
		pairs = append(pairs, wordFreqPair{word, freq})
	}
	
	// Sort by frequency (descending)
	for i := 0; i < len(pairs); i++ {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[i].freq < pairs[j].freq {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}
	
	// Take top N
	result := make(map[string]int)
	for i := 0; i < limit && i < len(pairs); i++ {
		result[pairs[i].word] = pairs[i].freq
	}
	
	return result
}

func handleWordCount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req WordCountRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if req.Text == "" {
		http.Error(w, "Text is required", http.StatusBadRequest)
		return
	}

	wordCount, wordFreq := countWords(req.Text)
	
	response := WordCountResponse{
		Text:          req.Text,
		WordCount:     wordCount,
		CharCount:     len(req.Text),
		FrequentWords: getTopWords(wordFreq, 5), // Show top 5 most frequent words
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/count", handleWordCount)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Word Counter function listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}