package errorx

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[REQUEST_ERROR] = "请求错误"
	message[REQUEST_PARAM_ERROR] = "参数错误"
	message[SERVER_COMMON_ERROR] = "服务器开小差啦,稍后再来试一试"
	message[TOKEN_EXPIRE_ERROR] = "令牌失效，请重新登陆"
	message[TOKEN_GENERATE_ERROR] = "生成令牌失败"
	message[DB_ERROR] = "系统繁忙,请稍后再试"
	message[DB_UPDATE_AFFECTED_ZERO_ERROR] = "更新数据失败,请稍后再试"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
