package internal

import (
	"encoding/json"
	"net/http"
)

type RestInfra struct {
	usecase Usecase
}

// NewRestInfra is responsible by building and returninig RestInfra.
func NewRestInfra(usecase Usecase) *RestInfra {
	return &RestInfra{usecase: usecase}
}

// GetByID returns a user by ID.
func (infra RestInfra) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, _ := infra.usecase.GetByID(1)
	json.NewEncoder(w).Encode(user)
}
