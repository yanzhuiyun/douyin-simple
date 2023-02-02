package result

import (
	"douyin-simple/common/xerr"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	StatusCode string `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

func (r *Response) structTomap() (map[string]interface{}, error) {
	out := make(map[string]interface{})
	data, _ := json.Marshal(r)
	json.Unmarshal(data, &out)
	return out, nil
}

func ResponseSuccess(w http.ResponseWriter, data any, tagName string) {
	reply := &Response{
		StatusCode: "0",
		StatusMsg:  "SUCCESS",
	}
	resp, _ := reply.structTomap()
	resp[tagName] = data
	httpx.OkJson(w, resp)
}

func ResponseBadRequest(w http.ResponseWriter, err error) {
	reply := &Response{
		StatusCode: xerr.ErrBadRequest,
		StatusMsg:  err.Error(),
	}
	httpx.WriteJson(w, http.StatusBadRequest, reply)
}

func ResponseBadServer(w http.ResponseWriter) {
	reply := &Response{
		StatusCode: xerr.ErrBadServer,
		StatusMsg:  xerr.MapErrMsg(xerr.ErrBadServer),
	}
	httpx.WriteJson(w, http.StatusServiceUnavailable, reply)
}

func GetTagName(ptr interface{}) string {
	if ptr == nil {
		return ""
	}
	typ := reflect.TypeOf(ptr).Elem()
	if typ.NumField() == 0 {
		return ""
	}
	return typ.Field(0).Tag.Get("json")
}
