package ascpffo

import (
	"sync"
)

// AliexpressAscpFroItemQueryData 结构体
type AliexpressAscpFroItemQueryData struct {
	// 商品Id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品名称
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 货品Id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 货品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 申请退货数量
	PlanQty string `json:"plan_qty,omitempty" xml:"plan_qty,omitempty"`
	// 价格
	OrderLinePrice string `json:"order_line_price,omitempty" xml:"order_line_price,omitempty"`
	// 仓库实收正品
	ReturnNormalQty string `json:"return_normal_qty,omitempty" xml:"return_normal_qty,omitempty"`
	// 仓库实收残品
	ReturnScrapQty string `json:"return_scrap_qty,omitempty" xml:"return_scrap_qty,omitempty"`
	// 仓名称
	WarehouseName string `json:"warehouse_name,omitempty" xml:"warehouse_name,omitempty"`
	// 扩展字段
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
}

var poolAliexpressAscpFroItemQueryData = sync.Pool{
	New: func() any {
		return new(AliexpressAscpFroItemQueryData)
	},
}

// GetAliexpressAscpFroItemQueryData() 从对象池中获取AliexpressAscpFroItemQueryData
func GetAliexpressAscpFroItemQueryData() *AliexpressAscpFroItemQueryData {
	return poolAliexpressAscpFroItemQueryData.Get().(*AliexpressAscpFroItemQueryData)
}

// ReleaseAliexpressAscpFroItemQueryData 释放AliexpressAscpFroItemQueryData
func ReleaseAliexpressAscpFroItemQueryData(v *AliexpressAscpFroItemQueryData) {
	v.ItemId = ""
	v.ItemTitle = ""
	v.ScItemId = ""
	v.Barcode = ""
	v.PlanQty = ""
	v.OrderLinePrice = ""
	v.ReturnNormalQty = ""
	v.ReturnScrapQty = ""
	v.WarehouseName = ""
	v.ExtendFields = ""
	poolAliexpressAscpFroItemQueryData.Put(v)
}
