package tmallcar

import (
	"sync"

	"github.com/bububa/opentaobao/model"
)

// ResultVo 结构体
type ResultVo struct {
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// costTime
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// gmtCurrentTime
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// 合同pdf字节码
	Objectbytes *model.File `json:"objectbytes,omitempty" xml:"objectbytes,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolResultVo = sync.Pool{
	New: func() any {
		return new(ResultVo)
	},
}

// GetResultVo() 从对象池中获取ResultVo
func GetResultVo() *ResultVo {
	return poolResultVo.Get().(*ResultVo)
}

// ReleaseResultVo 释放ResultVo
func ReleaseResultVo(v *ResultVo) {
	v.MsgCode = ""
	v.MsgInfo = ""
	v.CostTime = 0
	v.GmtCurrentTime = 0
	v.Objectbytes = nil
	v.Object = false
	v.Success = false
	poolResultVo.Put(v)
}
