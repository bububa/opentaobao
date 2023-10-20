package ascp

import (
	"sync"
)

// DataItem 结构体
type DataItem struct {
	// 组合货品erp货品id
	CombineScItemId string `json:"combine_sc_item_id,omitempty" xml:"combine_sc_item_id,omitempty"`
	// 错误码
	ErrCode string `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// 错误信息
	ErrMsg string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// ERP货品ID
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 卖家ID
	SellerId string `json:"seller_id,omitempty" xml:"seller_id,omitempty"`
	// 货品ID
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
	// 商品ID
	SkuId string `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// true|false
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolDataItem = sync.Pool{
	New: func() any {
		return new(DataItem)
	},
}

// GetDataItem() 从对象池中获取DataItem
func GetDataItem() *DataItem {
	return poolDataItem.Get().(*DataItem)
}

// ReleaseDataItem 释放DataItem
func ReleaseDataItem(v *DataItem) {
	v.CombineScItemId = ""
	v.ErrCode = ""
	v.ErrMsg = ""
	v.ScItemId = ""
	v.Code = ""
	v.Message = ""
	v.SellerId = ""
	v.ItemId = ""
	v.SkuId = ""
	v.Success = false
	poolDataItem.Put(v)
}
