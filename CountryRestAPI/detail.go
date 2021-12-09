package main

import (
	"encoding/json"
	"net/http"
)

// getCountry godoc
// @Summary Country Detail
// @Description Country Detail
// @Tags Country Detail
// @ID get-country
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string][]main.Country
// @Failure 400
// @Param Code path string true "Code"
// @Router /countries/ [get]
func getCountry(writer http.ResponseWriter, request *http.Request) {
	code := request.URL.Query()["Code"]
	country := data[code[0]]
	err := json.NewEncoder(writer).Encode(country)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
