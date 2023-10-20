package fenxiao

import (
	"sync"
)

// TipInfo 结构体
type TipInfo struct {
	// 商品id
	ScItemId string `json:"sc_item_id,omitempty" xml:"sc_item_id,omitempty"`
	// 返回信息
	Info string `json:"info,omitempty" xml:"info,omitempty"`
	// errorCode
	Errorcode string `json:"errorcode,omitempty" xml:"errorcode,omitempty"`
	// errorMessage
	Errormessage string `json:"errormessage,omitempty" xml:"errormessage,omitempty"`
	// scItemCode
	ScItemCode string `json:"sc_item_code,omitempty" xml:"sc_item_code,omitempty"`
	// storeCode
	StoreCode string `json:"store_code,omitempty" xml:"store_code,omitempty"`
	// invStoreType
	InvStoreType int64 `json:"inv_store_type,omitempty" xml:"inv_store_type,omitempty"`
	// skuId
	SkuId int64 `json:"sku_id,omitempty" xml:"sku_id,omitempty"`
	// 1前端商品 2供应链货品
	ItemType int64 `json:"item_type,omitempty" xml:"item_type,omitempty"`
}

var poolTipInfo = sync.Pool{
	New: func() any {
		return new(TipInfo)
	},
}

// GetTipInfo() 从对象池中获取TipInfo
func GetTipInfo() *TipInfo {
	return poolTipInfo.Get().(*TipInfo)
}

// ReleaseTipInfo 释放TipInfo
func ReleaseTipInfo(v *TipInfo) {
	v.ScItemId = ""
	v.Info = ""
	v.Errorcode = ""
	v.Errormessage = ""
	v.ScItemCode = ""
	v.StoreCode = ""
	v.InvStoreType = 0
	v.SkuId = 0
	v.ItemType = 0
	poolTipInfo.Put(v)
}
