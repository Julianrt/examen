package models

type Respuesta struct {
	Data       interface{} `json:"data"`
	CveError   int         `json:"cve_error"`
	CveMensaje string      `json:"cve_mensaje"`
}
