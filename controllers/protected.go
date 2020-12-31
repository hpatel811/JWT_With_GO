package controllers

import (
	"github.com/hpatel811/JWT_With_GO/models"
	"github.com/hpatel811/JWT_With_GO/utilities"
	"net/http"
)

func (c Controller) ProtectedEndPoint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("protectedEndpoint Invoked.")
		var errorMsg models.Error
		w.Header().Set("Content-Type", "application/JSON")
		w.WriteHeader(http.StatusOK)
		errorMsg.Message = "Successfully validated the protected endpoint with Token"
		utilities.ResponseJSON(w, errorMsg)
		return
	}
}
