package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// User error
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	ERROR_INVALID_TOKEN    = 1009

	// Article error
	ERROR_ART_NOT_EXIST = 2001
	ERROR_ART_IS_EMPTY  = 2002

	// Search error
	ERROR_KW_IS_EMPTY      = 3001
	ERROR_SQL_INJECTION    = 3002
	ERROR_RESULT_NOT_FOUND = 3003
)

var errormsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_INVALID_TOKEN:    "非法token",

	ERROR_ART_NOT_EXIST: "文章不存在",
	ERROR_ART_IS_EMPTY:  "文章内容不能为空！",

	ERROR_KW_IS_EMPTY:      "搜索关键词不能为空！",
	ERROR_SQL_INJECTION:    "非法输入！[SQL Injection]",
	ERROR_RESULT_NOT_FOUND: "没有相关搜索结果",
}

func GetErrorMessage(code int) string {
	return errormsg[code]
}
