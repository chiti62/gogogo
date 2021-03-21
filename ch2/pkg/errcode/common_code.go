package errcode

var (
	Success                   = NewError(0, "Success")
	ServerError               = NewError(10000000, "Server Error")
	InvalidParams             = NewError(10000001, "Invalid Parameters")
	NotFound                  = NewError(10000002, "Not Found")
	UnauthorizedAuthNotExist  = NewError(10000003, "Unauthorized, AppKey and AppSecret Not Found")
	UnauthorizedTokenError    = NewError(10000004, "Token Error")
	UnauthorizedTokenTimeout  = NewError(10000005, "Token Timeout")
	UnauthorizedTokenGenerate = NewError(10000006, "Failed to Generate Token")
	TooManyRequests           = NewError(10000007, "Too Many Requests")
)
