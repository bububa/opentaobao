package wdk

import (
	"sync"
)

// DelivererChangeRequest 结构体
type DelivererChangeRequest struct {
	// 经营店编码
	StoreId string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// 配送人员名称
	DelivererName string `json:"deliverer_name,omitempty" xml:"deliverer_name,omitempty"`
	// 配送人员联系方式
	DelivererPhone string `json:"deliverer_phone,omitempty" xml:"deliverer_phone,omitempty"`
	// 配送公司编码
	DelivererCompany string `json:"deliverer_company,omitempty" xml:"deliverer_company,omitempty"`
	// 配送物流单号
	LogisticsNo string `json:"logistics_no,omitempty" xml:"logistics_no,omitempty"`
	// 订单编码
	BizOrderId int64 `json:"biz_order_id,omitempty" xml:"biz_order_id,omitempty"`
}

var poolDelivererChangeRequest = sync.Pool{
	New: func() any {
		return new(DelivererChangeRequest)
	},
}

// GetDelivererChangeRequest() 从对象池中获取DelivererChangeRequest
func GetDelivererChangeRequest() *DelivererChangeRequest {
	return poolDelivererChangeRequest.Get().(*DelivererChangeRequest)
}

// ReleaseDelivererChangeRequest 释放DelivererChangeRequest
func ReleaseDelivererChangeRequest(v *DelivererChangeRequest) {
	v.StoreId = ""
	v.DelivererName = ""
	v.DelivererPhone = ""
	v.DelivererCompany = ""
	v.LogisticsNo = ""
	v.BizOrderId = 0
	poolDelivererChangeRequest.Put(v)
}
