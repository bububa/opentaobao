package wms

import (
	"sync"
)

// TaobaoWlbWmsSnInfoQueryResult 结构体
type TaobaoWlbWmsSnInfoQueryResult struct {
	// SN信息列表
	SnInfoList []Sninfolist `json:"sn_info_list,omitempty" xml:"sn_info_list>sninfolist,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTaobaoWlbWmsSnInfoQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoWlbWmsSnInfoQueryResult)
	},
}

// GetTaobaoWlbWmsSnInfoQueryResult() 从对象池中获取TaobaoWlbWmsSnInfoQueryResult
func GetTaobaoWlbWmsSnInfoQueryResult() *TaobaoWlbWmsSnInfoQueryResult {
	return poolTaobaoWlbWmsSnInfoQueryResult.Get().(*TaobaoWlbWmsSnInfoQueryResult)
}

// ReleaseTaobaoWlbWmsSnInfoQueryResult 释放TaobaoWlbWmsSnInfoQueryResult
func ReleaseTaobaoWlbWmsSnInfoQueryResult(v *TaobaoWlbWmsSnInfoQueryResult) {
	v.SnInfoList = v.SnInfoList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.TotalCount = 0
	v.Success = false
	poolTaobaoWlbWmsSnInfoQueryResult.Put(v)
}
