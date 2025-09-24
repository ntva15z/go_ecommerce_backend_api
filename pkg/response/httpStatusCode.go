package response

const (
	ErrCodeSuccess       = 20001 // success
	ErrCodeParamInvalid  = 20003 // email is invalid
	ErrInvalidToken      = 30001 // token invalid
	ErrCodeUserHasExists = 50001 // user has already registered
)

var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "email is invalid",
	ErrInvalidToken:      "token is invalid",
	ErrCodeUserHasExists: "user has already registered",
}
