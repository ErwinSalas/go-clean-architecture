package pkg

import "net/http"

type AccountHandler struct {
	service Service
}

func NewAccountHandler(service Service) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}
func (s *AccountHandler) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("Erwin", "salas")

	return writeJSON(w, http.StatusOK, account)
}

func (s *AccountHandler) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *AccountHandler) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
