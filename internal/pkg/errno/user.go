package errno

var (
	// ErrUserAlreadyExists 用户已存在
	ErrUserAlreadyExists = &Errno{Code: "FailedOperation.UserAlreadyExist", Message: "User already exist.", HTTP: 400}
)
