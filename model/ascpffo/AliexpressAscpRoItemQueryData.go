package ascpffo

import (
	"sync"
)

// AliexpressAscpRoItemQueryData 结构体
type AliexpressAscpRoItemQueryData struct {
	// 退供单号
	ReturnOrderNo string `json:"return_order_no,omitempty" xml:"return_order_no,omitempty"`
	// 货品名称
	Title string `json:"title,omitempty" xml:"title,omitempty"`
	// 库存类型
	InventoryTypeDesc string `json:"inventory_type_desc,omitempty" xml:"inventory_type_desc,omitempty"`
	// 税率
	TaxRate string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 含税退供价，单位分
	ReturnPrice string `json:"return_price,omitempty" xml:"return_price,omitempty"`
	// 扩展字段
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
	// 货品Id
	ScItemId int64 `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 退供数量
	Quantity int64 `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 实际退供数量
	ReturnQuantity int64 `json:"return_quantity,omitempty" xml:"return_quantity,omitempty"`
}

var poolAliexpressAscpRoItemQueryData = sync.Pool{
	New: func() any {
		return new(AliexpressAscpRoItemQueryData)
	},
}

// GetAliexpressAscpRoItemQueryData() 从对象池中获取AliexpressAscpRoItemQueryData
func GetAliexpressAscpRoItemQueryData() *AliexpressAscpRoItemQueryData {
	return poolAliexpressAscpRoItemQueryData.Get().(*AliexpressAscpRoItemQueryData)
}

// ReleaseAliexpressAscpRoItemQueryData 释放AliexpressAscpRoItemQueryData
func ReleaseAliexpressAscpRoItemQueryData(v *AliexpressAscpRoItemQueryData) {
	v.ReturnOrderNo = ""
	v.Title = ""
	v.InventoryTypeDesc = ""
	v.TaxRate = ""
	v.ReturnPrice = ""
	v.ExtendFields = ""
	v.ScItemId = 0
	v.Quantity = 0
	v.ReturnQuantity = 0
	poolAliexpressAscpRoItemQueryData.Put(v)
}
