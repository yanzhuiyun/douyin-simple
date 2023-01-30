package xerr

var message map[string]string

func init() {
	message = make(map[string]string)
	message[ErrBadRequest] = "请求报文存在语法错误或参数错误"
	message[ErrBadServer] = "服务端繁忙，稍后再试"
}

func MapErrMsg(code string) string {
	if msg, ok := message[code]; ok {
		return msg
	}
	return "Unknown Error"
}
