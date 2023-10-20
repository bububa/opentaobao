package ieagency

import (
	"sync"
)

// QueryChangeAgentListRs 结构体
type QueryChangeAgentListRs struct {
	// 改签单信息
	ChangeOrderDOs []IeChangeOrderVo `json:"change_order_d_os,omitempty" xml:"change_order_d_os>ie_change_order_vo,omitempty"`
	// 错误信息
	ApiErrorMsg string `json:"api_error_msg,omitempty" xml:"api_error_msg,omitempty"`
	// 错误码
	ApiErrorCode int64 `json:"api_error_code,omitempty" xml:"api_error_code,omitempty"`
	// 分页信息
	PageInfoDO *BasePageDo `json:"page_info_d_o,omitempty" xml:"page_info_d_o,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolQueryChangeAgentListRs = sync.Pool{
	New: func() any {
		return new(QueryChangeAgentListRs)
	},
}

// GetQueryChangeAgentListRs() 从对象池中获取QueryChangeAgentListRs
func GetQueryChangeAgentListRs() *QueryChangeAgentListRs {
	return poolQueryChangeAgentListRs.Get().(*QueryChangeAgentListRs)
}

// ReleaseQueryChangeAgentListRs 释放QueryChangeAgentListRs
func ReleaseQueryChangeAgentListRs(v *QueryChangeAgentListRs) {
	v.ChangeOrderDOs = v.ChangeOrderDOs[:0]
	v.ApiErrorMsg = ""
	v.ApiErrorCode = 0
	v.PageInfoDO = nil
	v.Success = false
	poolQueryChangeAgentListRs.Put(v)
}
