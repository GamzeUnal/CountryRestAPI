package main

import (
	"encoding/json"
	"net/http"
)

// getCountries godoc
// @Summary Country List
// @Description Country List
// @Tags Country List
// @ID get-countries
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]main.Country
// @Failure 400
// @Router /countries [get]
func getCountries(writer http.ResponseWriter, request *http.Request) {
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
