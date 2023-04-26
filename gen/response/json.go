package response

type JsonResponse struct {
	//Code    int         `json:"code"`
	Message string      `json:"msg"`
	Result  interface{} `json:"result"`
}

func New(message string, result interface{}) JsonResponse {
	return JsonResponse{
		//Code:    code,
		Message: message,
		Result:  result,
	}
}
