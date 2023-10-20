package ascpchannel

import (
	"sync"
)

// WaybillQueryRequestData 结构体
type WaybillQueryRequestData struct {
	// 商品信息
	Items []WaybillQueryPackageItem `json:"items,omitempty" xml:"items>waybill_query_package_item,omitempty"`
	// 包裹code，同一个订单唯一，同一个lp单号多次或同一次传重复的包裹号会被幂等
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 包裹描述信息
	Description string `json:"description,omitempty" xml:"description,omitempty"`
	// 子包裹数，在快运取号场景可能会用到
	TotalPackageCount string `json:"total_package_count,omitempty" xml:"total_package_count,omitempty"`
	// 包裹重量，单位克
	Weight string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 包裹体积，单位立方厘米
	Volume string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 长，单位mm
	Length string `json:"length,omitempty" xml:"length,omitempty"`
	// 宽，单位mm
	Width string `json:"width,omitempty" xml:"width,omitempty"`
	// 高，单位mm
	Height string `json:"height,omitempty" xml:"height,omitempty"`
}

var poolWaybillQueryRequestData = sync.Pool{
	New: func() any {
		return new(WaybillQueryRequestData)
	},
}

// GetWaybillQueryRequestData() 从对象池中获取WaybillQueryRequestData
func GetWaybillQueryRequestData() *WaybillQueryRequestData {
	return poolWaybillQueryRequestData.Get().(*WaybillQueryRequestData)
}

// ReleaseWaybillQueryRequestData 释放WaybillQueryRequestData
func ReleaseWaybillQueryRequestData(v *WaybillQueryRequestData) {
	v.Items = v.Items[:0]
	v.Code = ""
	v.Description = ""
	v.TotalPackageCount = ""
	v.Weight = ""
	v.Volume = ""
	v.Length = ""
	v.Width = ""
	v.Height = ""
	poolWaybillQueryRequestData.Put(v)
}
