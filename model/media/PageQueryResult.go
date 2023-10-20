package media

import (
	"sync"
)

// PageQueryResult 结构体
type PageQueryResult struct {
	// resultList
	ResultList []Resultlist `json:"result_list,omitempty" xml:"result_list>resultlist,omitempty"`
	// 总数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolPageQueryResult = sync.Pool{
	New: func() any {
		return new(PageQueryResult)
	},
}

// GetPageQueryResult() 从对象池中获取PageQueryResult
func GetPageQueryResult() *PageQueryResult {
	return poolPageQueryResult.Get().(*PageQueryResult)
}

// ReleasePageQueryResult 释放PageQueryResult
func ReleasePageQueryResult(v *PageQueryResult) {
	v.ResultList = v.ResultList[:0]
	v.Total = 0
	poolPageQueryResult.Put(v)
}
