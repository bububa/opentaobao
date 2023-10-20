package moscm

import (
	"sync"
)

// MosScrollQueryResult 结构体
type MosScrollQueryResult struct {
	// 返回数据
	Data []InventoryDetailDto `json:"data,omitempty" xml:"data>inventory_detail_dto,omitempty"`
	// 滚动查询id号
	ScrollId string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// 总条数
	Total int64 `json:"total,omitempty" xml:"total,omitempty"`
}

var poolMosScrollQueryResult = sync.Pool{
	New: func() any {
		return new(MosScrollQueryResult)
	},
}

// GetMosScrollQueryResult() 从对象池中获取MosScrollQueryResult
func GetMosScrollQueryResult() *MosScrollQueryResult {
	return poolMosScrollQueryResult.Get().(*MosScrollQueryResult)
}

// ReleaseMosScrollQueryResult 释放MosScrollQueryResult
func ReleaseMosScrollQueryResult(v *MosScrollQueryResult) {
	v.Data = v.Data[:0]
	v.ScrollId = ""
	v.Total = 0
	poolMosScrollQueryResult.Put(v)
}
