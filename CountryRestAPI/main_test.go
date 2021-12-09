package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostCountry(t *testing.T) {
	var jsonStr = []byte(`{"Code":"tr","Name":"Türkiye","ParentCity":"Ankara"}`)
	req, err := http.NewRequest("POST", "/newcountry/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	data = make(map[string]Country)
	handler := http.HandlerFunc(postCountry)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestListCountry(t *testing.T) {
	req, err := http.NewRequest("GET", "/countries", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(getCountries)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDetailCountry(t *testing.T) {
	req, err := http.NewRequest("GET", "/countries/?Code=tr", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(getCountry)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteCountry(t *testing.T) {
	req, err := http.NewRequest("GET", "countries/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(delCountries)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
