package logistic

import (
	"sync"
)

// WmsGoodsInfoSyncRequest 结构体
type WmsGoodsInfoSyncRequest struct {
	// 货主code
	WmsOwnerCode string `json:"wms_owner_code,omitempty" xml:"wms_owner_code,omitempty"`
	// 仓code
	WmsWarehouseCode string `json:"wms_warehouse_code,omitempty" xml:"wms_warehouse_code,omitempty"`
	// 货品code
	ItemCode string `json:"item_code,omitempty" xml:"item_code,omitempty"`
	// 长 (厘米)
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	//  宽 (厘米)
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 高 (厘米)
	Height string `json:"height,omitempty" xml:"height,omitempty"`
	// 体积 (升)
	Volume string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 毛重 (千克)
	GrossWeight string `json:"gross_weight,omitempty" xml:"gross_weight,omitempty"`
	// 净重 (千克)
	NetWeight string `json:"net_weight,omitempty" xml:"net_weight,omitempty"`
	// 图片链接
	Pic string `json:"pic,omitempty" xml:"pic,omitempty"`
	// 扩展字段
	Feature string `json:"feature,omitempty" xml:"feature,omitempty"`
	// 货品id
	ItemId string `json:"item_id,omitempty" xml:"item_id,omitempty"`
}

var poolWmsGoodsInfoSyncRequest = sync.Pool{
	New: func() any {
		return new(WmsGoodsInfoSyncRequest)
	},
}

// GetWmsGoodsInfoSyncRequest() 从对象池中获取WmsGoodsInfoSyncRequest
func GetWmsGoodsInfoSyncRequest() *WmsGoodsInfoSyncRequest {
	return poolWmsGoodsInfoSyncRequest.Get().(*WmsGoodsInfoSyncRequest)
}

// ReleaseWmsGoodsInfoSyncRequest 释放WmsGoodsInfoSyncRequest
func ReleaseWmsGoodsInfoSyncRequest(v *WmsGoodsInfoSyncRequest) {
	v.WmsOwnerCode = ""
	v.WmsWarehouseCode = ""
	v.ItemCode = ""
	v.Length = ""
	v.Width = ""
	v.Height = ""
	v.Volume = ""
	v.GrossWeight = ""
	v.NetWeight = ""
	v.Pic = ""
	v.Feature = ""
	v.ItemId = ""
	poolWmsGoodsInfoSyncRequest.Put(v)
}
