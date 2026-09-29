package main

type Field struct {
	ID        *int    `json:"ID"`
	Title     *string `json:"Title"`
	Completed *bool   `json:"Completed"`
}

type Temp struct {
	Title     *string
	Completed *bool
}

type ErrResponse struct {
	Error string
}
