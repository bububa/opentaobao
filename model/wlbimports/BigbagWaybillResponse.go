package wlbimports

import (
	"sync"
)

// BigbagWaybillResponse 结构体
type BigbagWaybillResponse struct {
	// 大包Code
	BigbagCode string `json:"bigbag_code,omitempty" xml:"bigbag_code,omitempty"`
	// 大包运单号
	TrackingNumber string `json:"tracking_number,omitempty" xml:"tracking_number,omitempty"`
	// 大包子运单号集合（用,分隔 ）
	SubWaybillNos string `json:"sub_waybill_nos,omitempty" xml:"sub_waybill_nos,omitempty"`
	// 面单内荣
	FileContext string `json:"file_context,omitempty" xml:"file_context,omitempty"`
	// 面单文件名称
	FileName string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// 面单格式
	ImageFormat string `json:"image_format,omitempty" xml:"image_format,omitempty"`
	// 大包id
	BigbagId int64 `json:"bigbag_id,omitempty" xml:"bigbag_id,omitempty"`
}

var poolBigbagWaybillResponse = sync.Pool{
	New: func() any {
		return new(BigbagWaybillResponse)
	},
}

// GetBigbagWaybillResponse() 从对象池中获取BigbagWaybillResponse
func GetBigbagWaybillResponse() *BigbagWaybillResponse {
	return poolBigbagWaybillResponse.Get().(*BigbagWaybillResponse)
}

// ReleaseBigbagWaybillResponse 释放BigbagWaybillResponse
func ReleaseBigbagWaybillResponse(v *BigbagWaybillResponse) {
	v.BigbagCode = ""
	v.TrackingNumber = ""
	v.SubWaybillNos = ""
	v.FileContext = ""
	v.FileName = ""
	v.ImageFormat = ""
	v.BigbagId = 0
	poolBigbagWaybillResponse.Put(v)
}
