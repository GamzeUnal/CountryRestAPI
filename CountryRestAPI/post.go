package main

import (
	"encoding/json"
	"net/http"
)

// postCountry godoc
// @Summary Post Country
// @Description Create Country
// @Tags post
// @ID post-country
// @Accept  json
// @Produce  json
// @Param postCountry body Country true "Post Country Input"
// @Success 200 {object} Country
// @Failure 400
// @Router /newcountry [post]
func postCountry(writer http.ResponseWriter, request *http.Request) {
	var country Country
	err := json.NewDecoder(request.Body).Decode(&country)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	data[country.Code] = country
	err = json.NewEncoder(writer).Encode(country)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}
