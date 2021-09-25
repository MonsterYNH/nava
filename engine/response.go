package engine

type APIResponse struct {
	ErrCode
	Data interface{} `json:"data"`
}

func NewApiResponse(data interface{}, err error) APIResponse {
	switch errCode := err.(type) {
	case ErrCode:
		return APIResponse{
			ErrCode: errCode,
			Data:    data,
		}
	default:
		return APIResponse{
			ErrCode: ErrUnknown,
		}
	}
}
