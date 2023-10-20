package tuike

import (
	"sync"
)

// TaOfferSearchResult 结构体
type TaOfferSearchResult struct {
	// 数据
	DataList []string `json:"data_list,omitempty" xml:"data_list>string,omitempty"`
	// 错误信息
	Errors string `json:"errors,omitempty" xml:"errors,omitempty"`
	// 请求状态
	Status string `json:"status,omitempty" xml:"status,omitempty"`
	// 查询总记录数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
	// 当前条数
	Num int64 `json:"num,omitempty" xml:"num,omitempty"`
	// 分页大小
	PageSize int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 当前页
	PageNum int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
}

var poolTaOfferSearchResult = sync.Pool{
	New: func() any {
		return new(TaOfferSearchResult)
	},
}

// GetTaOfferSearchResult() 从对象池中获取TaOfferSearchResult
func GetTaOfferSearchResult() *TaOfferSearchResult {
	return poolTaOfferSearchResult.Get().(*TaOfferSearchResult)
}

// ReleaseTaOfferSearchResult 释放TaOfferSearchResult
func ReleaseTaOfferSearchResult(v *TaOfferSearchResult) {
	v.DataList = v.DataList[:0]
	v.Errors = ""
	v.Status = ""
	v.Total = 0
	v.Num = 0
	v.PageSize = 0
	v.PageNum = 0
	poolTaOfferSearchResult.Put(v)
}
