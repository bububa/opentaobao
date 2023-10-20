package wdk

import (
	"sync"
)

// InventoryTopResultBo 结构体
type InventoryTopResultBo struct {
	// inventoryInfoList
	InventoryInfoList []Inventoryinfolist `json:"inventory_info_list,omitempty" xml:"inventory_info_list>inventoryinfolist,omitempty"`
	// queryPageNum
	QueryPageNum int64 `json:"query_page_num,omitempty" xml:"query_page_num,omitempty"`
	// queryPageSize
	QueryPageSize int64 `json:"query_page_size,omitempty" xml:"query_page_size,omitempty"`
	// totalPageNum
	TotalPageNum int64 `json:"total_page_num,omitempty" xml:"total_page_num,omitempty"`
}

var poolInventoryTopResultBo = sync.Pool{
	New: func() any {
		return new(InventoryTopResultBo)
	},
}

// GetInventoryTopResultBo() 从对象池中获取InventoryTopResultBo
func GetInventoryTopResultBo() *InventoryTopResultBo {
	return poolInventoryTopResultBo.Get().(*InventoryTopResultBo)
}

// ReleaseInventoryTopResultBo 释放InventoryTopResultBo
func ReleaseInventoryTopResultBo(v *InventoryTopResultBo) {
	v.InventoryInfoList = v.InventoryInfoList[:0]
	v.QueryPageNum = 0
	v.QueryPageSize = 0
	v.TotalPageNum = 0
	poolInventoryTopResultBo.Put(v)
}
