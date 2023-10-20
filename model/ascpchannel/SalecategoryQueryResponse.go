package ascpchannel

import (
	"sync"
)

// SalecategoryQueryResponse 结构体
type SalecategoryQueryResponse struct {
	// 货品品类名称
	ItemSalecategoryName string `json:"item_salecategory_name,omitempty" xml:"item_salecategory_name,omitempty"`
	// 货品ID
	ItemId int64 `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 货品品类ID
	ItemSalecategory int64 `json:"item_salecategory,omitempty" xml:"item_salecategory,omitempty"`
}

var poolSalecategoryQueryResponse = sync.Pool{
	New: func() any {
		return new(SalecategoryQueryResponse)
	},
}

// GetSalecategoryQueryResponse() 从对象池中获取SalecategoryQueryResponse
func GetSalecategoryQueryResponse() *SalecategoryQueryResponse {
	return poolSalecategoryQueryResponse.Get().(*SalecategoryQueryResponse)
}

// ReleaseSalecategoryQueryResponse 释放SalecategoryQueryResponse
func ReleaseSalecategoryQueryResponse(v *SalecategoryQueryResponse) {
	v.ItemSalecategoryName = ""
	v.ItemId = 0
	v.ItemSalecategory = 0
	poolSalecategoryQueryResponse.Put(v)
}
