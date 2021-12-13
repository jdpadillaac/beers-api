package controllers

const (
	requestSuccessMsg  = "Preteición servida de manera correcta"
	badDataMsg         = "Datos inconrectos"
	registerSuccessMsg = "Registro exitoso"
	registerFailedMsg  = "Registro fallido"
)

type ctrResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func newResponse(status bool, message string, data interface{}) ctrResponse {
	return ctrResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
