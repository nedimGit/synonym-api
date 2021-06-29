package endpoints

import (
	"io"
	"net/http"

	m "github.com/NedimUka/synonyms/endpoints/middleware"
	"github.com/NedimUka/synonyms/endpoints/synonym"
)

// HealthCheck method is used to check if server is live.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "OK")
}

// Initialize is used to setup
func Initialize() {
	http.HandleFunc("/", HealthCheck)

	// Add new word
	http.HandleFunc("/synonym/word/add", m.Chain(m.InitMiddleware, m.Post).Then(synonym.AddWord))
	// Get all words in system
	http.HandleFunc("/synonym/word/list", m.Chain(m.InitMiddleware, m.Get).Then(synonym.GetWords))
	// Search words
	http.HandleFunc("/synonym/word/search", m.Chain(m.InitMiddleware, m.Post).Then(synonym.SearchWords))
	// Add snym to existing word
	http.HandleFunc("/synonym/word/synonym/add", m.Chain(m.InitMiddleware, m.Post).Then(synonym.AddSynonym))
	// Update synonym
	http.HandleFunc("/synonym/word/synonym/update", m.Chain(m.InitMiddleware, m.Post).Then(synonym.UpdateSynonym))
	// Remove synonym
	http.HandleFunc("/synonym/word/synonym/remove", m.Chain(m.InitMiddleware, m.Post).Then(synonym.RemoveSynonym))
	// List all synonyms of specific word
	http.HandleFunc("/synonym/word/synonym/list", m.Chain(m.InitMiddleware, m.Post).Then(synonym.GetSynonyms))

}
