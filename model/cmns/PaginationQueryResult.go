package cmns

import (
	"sync"
)

// PaginationQueryResult 结构体
type PaginationQueryResult struct {
	// ack记录列表
	List []MessageAckResult `json:"list,omitempty" xml:"list>message_ack_result,omitempty"`
	// 分页数据
	Pagination *Pagination `json:"pagination,omitempty" xml:"pagination,omitempty"`
}

var poolPaginationQueryResult = sync.Pool{
	New: func() any {
		return new(PaginationQueryResult)
	},
}

// GetPaginationQueryResult() 从对象池中获取PaginationQueryResult
func GetPaginationQueryResult() *PaginationQueryResult {
	return poolPaginationQueryResult.Get().(*PaginationQueryResult)
}

// ReleasePaginationQueryResult 释放PaginationQueryResult
func ReleasePaginationQueryResult(v *PaginationQueryResult) {
	v.List = v.List[:0]
	v.Pagination = nil
	poolPaginationQueryResult.Put(v)
}
