package models

type Ressponse struct {
	Success bool   `json:"success"`
	Massage string `json:"massage"`
	Data    any    `json:"data,omitempty"`
}