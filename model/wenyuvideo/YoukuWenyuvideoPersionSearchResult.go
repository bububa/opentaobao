package wenyuvideo

import (
	"sync"
)

// YoukuWenyuvideoPersionSearchResult 结构体
type YoukuWenyuvideoPersionSearchResult struct {
	// 业务扩展数据
	BizExtMap string `json:"biz_ext_map,omitempty" xml:"biz_ext_map,omitempty"`
	// 业务错误码
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 业务错误提示
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 人物列表
	Model *PersonDo `json:"model,omitempty" xml:"model,omitempty"`
	// HTTP请求状态
	HttpStatusCode int64 `json:"http_status_code,omitempty" xml:"http_status_code,omitempty"`
	// 接口调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolYoukuWenyuvideoPersionSearchResult = sync.Pool{
	New: func() any {
		return new(YoukuWenyuvideoPersionSearchResult)
	},
}

// GetYoukuWenyuvideoPersionSearchResult() 从对象池中获取YoukuWenyuvideoPersionSearchResult
func GetYoukuWenyuvideoPersionSearchResult() *YoukuWenyuvideoPersionSearchResult {
	return poolYoukuWenyuvideoPersionSearchResult.Get().(*YoukuWenyuvideoPersionSearchResult)
}

// ReleaseYoukuWenyuvideoPersionSearchResult 释放YoukuWenyuvideoPersionSearchResult
func ReleaseYoukuWenyuvideoPersionSearchResult(v *YoukuWenyuvideoPersionSearchResult) {
	v.BizExtMap = ""
	v.MsgCode = ""
	v.MsgInfo = ""
	v.Model = nil
	v.HttpStatusCode = 0
	v.Success = false
	poolYoukuWenyuvideoPersionSearchResult.Put(v)
}
