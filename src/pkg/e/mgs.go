package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR : "fail",
	INVAILD_PARAMS : "请求参数错误",
	ERROR_EXIST_TAG  : "已存在该标签",
	ERROR_NOT_EXIST_TAG : "该标签不存在",
	ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token 验证失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token 验证超时",
	ERROR_AUTH_TOKEN : "Token 生成失败",
	ERROR_AUTH : "Token 错误",
}

func  GetMsg(code int) string  {
	msg,ok:= MsgFlags[code]
	if ok{
		return  msg
	}else {
		return  MsgFlags[ERROR]
	}
}