package main

type Country struct {
	Code       string `json:"Code"`
	Name       string `json:"Name"`
	ParentCity string `json:"ParentCity"`
}

// swagger:parameters getCountry
type Params struct {
	// UserId that identifies a user.
	//
	// in: query
	Code string `json:"Code"`
}
