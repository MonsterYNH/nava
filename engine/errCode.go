package engine

type ErrCode struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errs    []string `json:"errors"`
}

func (e ErrCode) Error() string {
	return e.Message
}

func (e *ErrCode) SetErrors(errs ...string) ErrCode {
	if len(e.Errs) == 0 {
		e.Errs = []string{}
	}
	return ErrCode{
		Code:    e.Code,
		Message: e.Message,
		Errs:    append(e.Errs, errs...),
	}
}

func NewErrors(code int, message string, errs ...string) ErrCode {
	return ErrCode{
		Code:    code,
		Message: message,
		Errs:    errs,
	}
}

var (
	ErrSuccess = ErrCode{Code: 0, Message: "Success", Errs: []string{}}
	ErrUnknown = ErrCode{Code: 100001, Message: "Unknown error, please contact the developer", Errs: []string{}}
	ErrNoAuth  = ErrCode{Code: 100002, Message: "Authorized failed", Errs: []string{}}

	ErrParam = ErrCode{Code: 100101, Message: "Param error", Errs: []string{}}
)
