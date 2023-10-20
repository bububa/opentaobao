package logistic

import (
	"sync"
)

// Tms2MscPayQueryRequest 结构体
type Tms2MscPayQueryRequest struct {
	// 服务类型
	ServiceType string `json:"service_type,omitempty" xml:"service_type,omitempty"`
	// 服务商ID
	SupplierId string `json:"supplier_id,omitempty" xml:"supplier_id,omitempty"`
	// 配资源编码
	BizCode string `json:"biz_code,omitempty" xml:"biz_code,omitempty"`
	// 唯一标识单号
	DeliveryCode string `json:"delivery_code,omitempty" xml:"delivery_code,omitempty"`
}

var poolTms2MscPayQueryRequest = sync.Pool{
	New: func() any {
		return new(Tms2MscPayQueryRequest)
	},
}

// GetTms2MscPayQueryRequest() 从对象池中获取Tms2MscPayQueryRequest
func GetTms2MscPayQueryRequest() *Tms2MscPayQueryRequest {
	return poolTms2MscPayQueryRequest.Get().(*Tms2MscPayQueryRequest)
}

// ReleaseTms2MscPayQueryRequest 释放Tms2MscPayQueryRequest
func ReleaseTms2MscPayQueryRequest(v *Tms2MscPayQueryRequest) {
	v.ServiceType = ""
	v.SupplierId = ""
	v.BizCode = ""
	v.DeliveryCode = ""
	poolTms2MscPayQueryRequest.Put(v)
}
