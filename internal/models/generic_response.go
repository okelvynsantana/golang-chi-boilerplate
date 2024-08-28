package models

type GenericResponse struct {
	Message string `json:"message"`
	Status  string `json:"status_code"`
	Code    int    `json:"code"`
}
