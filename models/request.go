package models

type Request struct {
	Title string `json:"title"`
}

type RequestPoint struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}