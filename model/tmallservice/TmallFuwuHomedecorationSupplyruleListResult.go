package tmallservice

import (
	"sync"
)

// TmallFuwuHomedecorationSupplyruleListResult 结构体
type TmallFuwuHomedecorationSupplyruleListResult struct {
	// 规则数据
	DataList []SupplyConfigRuleDto `json:"data_list,omitempty" xml:"data_list>supply_config_rule_dto,omitempty"`
	// 错误码
	ErrorCode string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// 错误信息
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 当前页码
	PageIndex int64 `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 每页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总页数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 是否成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolTmallFuwuHomedecorationSupplyruleListResult = sync.Pool{
	New: func() any {
		return new(TmallFuwuHomedecorationSupplyruleListResult)
	},
}

// GetTmallFuwuHomedecorationSupplyruleListResult() 从对象池中获取TmallFuwuHomedecorationSupplyruleListResult
func GetTmallFuwuHomedecorationSupplyruleListResult() *TmallFuwuHomedecorationSupplyruleListResult {
	return poolTmallFuwuHomedecorationSupplyruleListResult.Get().(*TmallFuwuHomedecorationSupplyruleListResult)
}

// ReleaseTmallFuwuHomedecorationSupplyruleListResult 释放TmallFuwuHomedecorationSupplyruleListResult
func ReleaseTmallFuwuHomedecorationSupplyruleListResult(v *TmallFuwuHomedecorationSupplyruleListResult) {
	v.DataList = v.DataList[:0]
	v.ErrorCode = ""
	v.ErrorMsg = ""
	v.PageIndex = 0
	v.PageSize = 0
	v.TotalCount = 0
	v.Success = false
	poolTmallFuwuHomedecorationSupplyruleListResult.Put(v)
}
