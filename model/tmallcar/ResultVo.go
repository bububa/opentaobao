package tmallcar

import (
	"github.com/bububa/opentaobao/model"
)

// ResultVo 结构体
type ResultVo struct {
	// costTime
	CostTime int64 `json:"cost_time,omitempty" xml:"cost_time,omitempty"`
	// gmtCurrentTime
	GmtCurrentTime int64 `json:"gmt_current_time,omitempty" xml:"gmt_current_time,omitempty"`
	// msgCode
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// msgInfo
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// object
	Object bool `json:"object,omitempty" xml:"object,omitempty"`
	// success
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
	// 合同pdf字节码
	Objectbytes *model.File `json:"objectbytes,omitempty" xml:"objectbytes,omitempty"`
}
