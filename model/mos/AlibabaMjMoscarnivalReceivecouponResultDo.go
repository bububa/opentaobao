package mos

import (
	"sync"
)

// AlibabaMjMoscarnivalReceivecouponResultDo 结构体
type AlibabaMjMoscarnivalReceivecouponResultDo struct {
	// 标题
	Titles []string `json:"titles,omitempty" xml:"titles>string,omitempty"`
	// 调用链id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 是否成功
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 总行数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMjMoscarnivalReceivecouponResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMjMoscarnivalReceivecouponResultDo)
	},
}

// GetAlibabaMjMoscarnivalReceivecouponResultDo() 从对象池中获取AlibabaMjMoscarnivalReceivecouponResultDo
func GetAlibabaMjMoscarnivalReceivecouponResultDo() *AlibabaMjMoscarnivalReceivecouponResultDo {
	return poolAlibabaMjMoscarnivalReceivecouponResultDo.Get().(*AlibabaMjMoscarnivalReceivecouponResultDo)
}

// ReleaseAlibabaMjMoscarnivalReceivecouponResultDo 释放AlibabaMjMoscarnivalReceivecouponResultDo
func ReleaseAlibabaMjMoscarnivalReceivecouponResultDo(v *AlibabaMjMoscarnivalReceivecouponResultDo) {
	v.Titles = v.Titles[:0]
	v.TraceId = ""
	v.Data = ""
	v.ErrMsg = ""
	v.ResultCode = ""
	v.Total = 0
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMjMoscarnivalReceivecouponResultDo.Put(v)
}
