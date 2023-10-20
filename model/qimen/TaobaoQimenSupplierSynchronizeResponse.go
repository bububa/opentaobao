package qimen

import (
	"sync"
)

// TaobaoQimenSupplierSynchronizeResponse 结构体
type TaobaoQimenSupplierSynchronizeResponse struct {
	// success|failure
	Flag string `json:"flag,omitempty" xml:"flag,omitempty"`
	// 响应码
	Code string `json:"code,omitempty" xml:"code,omitempty"`
	// 响应信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
}

var poolTaobaoQimenSupplierSynchronizeResponse = sync.Pool{
	New: func() any {
		return new(TaobaoQimenSupplierSynchronizeResponse)
	},
}

// GetTaobaoQimenSupplierSynchronizeResponse() 从对象池中获取TaobaoQimenSupplierSynchronizeResponse
func GetTaobaoQimenSupplierSynchronizeResponse() *TaobaoQimenSupplierSynchronizeResponse {
	return poolTaobaoQimenSupplierSynchronizeResponse.Get().(*TaobaoQimenSupplierSynchronizeResponse)
}

// ReleaseTaobaoQimenSupplierSynchronizeResponse 释放TaobaoQimenSupplierSynchronizeResponse
func ReleaseTaobaoQimenSupplierSynchronizeResponse(v *TaobaoQimenSupplierSynchronizeResponse) {
	v.Flag = ""
	v.Code = ""
	v.Message = ""
	poolTaobaoQimenSupplierSynchronizeResponse.Put(v)
}
