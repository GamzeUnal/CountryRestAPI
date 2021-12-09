package main

import (
	"net/http"
)

// delCountries godoc
// @Summary Remove Countries
// @Description Remove Countries
// @Tags Country Delete
// @ID del-countries
// @Accept  json
// @Produce  json
// @Success 200
// @Router /countries/delete [delete]
func delCountries(writer http.ResponseWriter, request *http.Request) {
	data = make(map[string]Country)
}
