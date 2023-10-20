package fenxiao

import (
	"sync"
)

// PduList 结构体
type PduList struct {
	// 分销商用户名
	DistributorName string `json:"distributor_name,omitempty" xml:"distributor_name,omitempty"`
	// 产品销售属性值
	SkuProperties string `json:"sku_properties,omitempty" xml:"sku_properties,omitempty"`
	// 产品ID
	ProductId int64 `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 分销商ID
	DistributorId int64 `json:"distributor_id,omitempty" xml:"distributor_id,omitempty"`
	// 产品代销配额库存
	QuantityAgent int64 `json:"quantity_agent,omitempty" xml:"quantity_agent,omitempty"`
}

var poolPduList = sync.Pool{
	New: func() any {
		return new(PduList)
	},
}

// GetPduList() 从对象池中获取PduList
func GetPduList() *PduList {
	return poolPduList.Get().(*PduList)
}

// ReleasePduList 释放PduList
func ReleasePduList(v *PduList) {
	v.DistributorName = ""
	v.SkuProperties = ""
	v.ProductId = 0
	v.DistributorId = 0
	v.QuantityAgent = 0
	poolPduList.Put(v)
}
