package ascpffo

import (
	"sync"
)

// AliexpressAscpFfoItemQueryData 结构体
type AliexpressAscpFfoItemQueryData struct {
	// 库存数量
	AicInventory string `json:"aic_inventory,omitempty" xml:"aic_inventory,omitempty"`
	// 货品实际支付金额
	SkuActualPaidAmount string `json:"sku_actual_paid_amount,omitempty" xml:"sku_actual_paid_amount,omitempty"`
	// 货品优惠金额
	SkuDiscountAmount string `json:"sku_discount_amount,omitempty" xml:"sku_discount_amount,omitempty"`
	// SKU单价
	UnitPrice string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
	// 发货数量
	OrderLineQty string `json:"order_line_qty,omitempty" xml:"order_line_qty,omitempty"`
	// 货品条码
	Barcode string `json:"barcode,omitempty" xml:"barcode,omitempty"`
	// 货品Id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// SKUid
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 商品名称
	ItemTitle string `json:"item_title,omitempty" xml:"item_title,omitempty"`
	// 商品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 扩展字段
	ExtendFields string `json:"extend_fields,omitempty" xml:"extend_fields,omitempty"`
}

var poolAliexpressAscpFfoItemQueryData = sync.Pool{
	New: func() any {
		return new(AliexpressAscpFfoItemQueryData)
	},
}

// GetAliexpressAscpFfoItemQueryData() 从对象池中获取AliexpressAscpFfoItemQueryData
func GetAliexpressAscpFfoItemQueryData() *AliexpressAscpFfoItemQueryData {
	return poolAliexpressAscpFfoItemQueryData.Get().(*AliexpressAscpFfoItemQueryData)
}

// ReleaseAliexpressAscpFfoItemQueryData 释放AliexpressAscpFfoItemQueryData
func ReleaseAliexpressAscpFfoItemQueryData(v *AliexpressAscpFfoItemQueryData) {
	v.AicInventory = ""
	v.SkuActualPaidAmount = ""
	v.SkuDiscountAmount = ""
	v.UnitPrice = ""
	v.OrderLineQty = ""
	v.Barcode = ""
	v.ScItemId = ""
	v.SkuId = ""
	v.ItemTitle = ""
	v.ItemId = ""
	v.ExtendFields = ""
	poolAliexpressAscpFfoItemQueryData.Put(v)
}
