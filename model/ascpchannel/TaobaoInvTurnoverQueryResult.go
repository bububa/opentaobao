package ascpchannel

import (
	"sync"
)

// TaobaoInvTurnoverQueryResult 结构体
type TaobaoInvTurnoverQueryResult struct {
	// 流水信息&lt;明细&gt;
	DataList []TaobaoInvTurnoverQueryData `json:"data_list,omitempty" xml:"data_list>taobao_inv_turnover_query_data,omitempty"`
	// 查询页
	PageIndex string `json:"page_index,omitempty" xml:"page_index,omitempty"`
	// 单页记录数
	PageSize string `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总记录数
	TotalCount string `json:"total_count,omitempty" xml:"total_count,omitempty"`
	// 当前页数
	CurrentPageCount string `json:"current_page_count,omitempty" xml:"current_page_count,omitempty"`
	// 总页数
	TotalPage string `json:"total_page,omitempty" xml:"total_page,omitempty"`
}

var poolTaobaoInvTurnoverQueryResult = sync.Pool{
	New: func() any {
		return new(TaobaoInvTurnoverQueryResult)
	},
}

// GetTaobaoInvTurnoverQueryResult() 从对象池中获取TaobaoInvTurnoverQueryResult
func GetTaobaoInvTurnoverQueryResult() *TaobaoInvTurnoverQueryResult {
	return poolTaobaoInvTurnoverQueryResult.Get().(*TaobaoInvTurnoverQueryResult)
}

// ReleaseTaobaoInvTurnoverQueryResult 释放TaobaoInvTurnoverQueryResult
func ReleaseTaobaoInvTurnoverQueryResult(v *TaobaoInvTurnoverQueryResult) {
	v.DataList = v.DataList[:0]
	v.PageIndex = ""
	v.PageSize = ""
	v.TotalCount = ""
	v.CurrentPageCount = ""
	v.TotalPage = ""
	poolTaobaoInvTurnoverQueryResult.Put(v)
}
