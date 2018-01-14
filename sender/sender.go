package sender

import (
	"encoding/json"
	"net/http"

	"github.com/GnuYtuce/ytuyemekhane-api/models"
)

// JSON :
func JSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	jContent, _ := json.MarshalIndent(data, "", " ")
	w.Write(jContent)
}

// Err :
func Err(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	jContent, _ := json.MarshalIndent(models.Error{Content: err.Error()}, "", " ")
	w.Write(jContent)
}
