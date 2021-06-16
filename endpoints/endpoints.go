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

	// Used to add new word with synonym
	http.HandleFunc("/synonym/word/add", m.Chain(m.Post).Then(synonym.AddWord))

}
