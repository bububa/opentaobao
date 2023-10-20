package mos

import (
	"sync"
)

// AlibabaMosBunkBunkinfoQuerybunkResultDo 结构体
type AlibabaMosBunkBunkinfoQuerybunkResultDo struct {
	// 返回数据
	DataList []BunkSimpleDto `json:"data_list,omitempty" xml:"data_list>bunk_simple_dto,omitempty"`
	// 结果标题
	Titles []string `json:"titles,omitempty" xml:"titles>string,omitempty"`
	// 全链路追踪id
	TraceId string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
	// 其他
	Extra string `json:"extra,omitempty" xml:"extra,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 结果码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 总量
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 错误码
	ErrCode int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 调用是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaMosBunkBunkinfoQuerybunkResultDo = sync.Pool{
	New: func() any {
		return new(AlibabaMosBunkBunkinfoQuerybunkResultDo)
	},
}

// GetAlibabaMosBunkBunkinfoQuerybunkResultDo() 从对象池中获取AlibabaMosBunkBunkinfoQuerybunkResultDo
func GetAlibabaMosBunkBunkinfoQuerybunkResultDo() *AlibabaMosBunkBunkinfoQuerybunkResultDo {
	return poolAlibabaMosBunkBunkinfoQuerybunkResultDo.Get().(*AlibabaMosBunkBunkinfoQuerybunkResultDo)
}

// ReleaseAlibabaMosBunkBunkinfoQuerybunkResultDo 释放AlibabaMosBunkBunkinfoQuerybunkResultDo
func ReleaseAlibabaMosBunkBunkinfoQuerybunkResultDo(v *AlibabaMosBunkBunkinfoQuerybunkResultDo) {
	v.DataList = v.DataList[:0]
	v.Titles = v.Titles[:0]
	v.TraceId = ""
	v.Extra = ""
	v.ErrMsg = ""
	v.ResultCode = ""
	v.Total = 0
	v.ErrCode = 0
	v.Success = false
	poolAlibabaMosBunkBunkinfoQuerybunkResultDo.Put(v)
}
