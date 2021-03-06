package handler

import (
	"net/http"

	"github.com/achanda/testrest/version"
	"github.com/jinzhu/gorm"
)

func GetVersion(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, version.Version)
}
