package wdk

import (
	"sync"
)

// WmsInventoryTopQuery 结构体
type WmsInventoryTopQuery struct {
	// 仓库编号
	WarehouseCode string `json:"warehouse_code,omitempty" xml:"warehouse_code,omitempty"`
	// 查询页数
	QueryPageNum int64 `json:"query_page_num,omitempty" xml:"query_page_num,omitempty"`
}

var poolWmsInventoryTopQuery = sync.Pool{
	New: func() any {
		return new(WmsInventoryTopQuery)
	},
}

// GetWmsInventoryTopQuery() 从对象池中获取WmsInventoryTopQuery
func GetWmsInventoryTopQuery() *WmsInventoryTopQuery {
	return poolWmsInventoryTopQuery.Get().(*WmsInventoryTopQuery)
}

// ReleaseWmsInventoryTopQuery 释放WmsInventoryTopQuery
func ReleaseWmsInventoryTopQuery(v *WmsInventoryTopQuery) {
	v.WarehouseCode = ""
	v.QueryPageNum = 0
	poolWmsInventoryTopQuery.Put(v)
}
