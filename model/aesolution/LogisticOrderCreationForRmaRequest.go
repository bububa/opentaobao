package aesolution

import (
	"sync"
)

// LogisticOrderCreationForRmaRequest 结构体
type LogisticOrderCreationForRmaRequest struct {
	// Carrier tracking code. Tracking code or Shipping code must be provided
	TrackingCode string `json:"tracking_code,omitempty" xml:"tracking_code,omitempty"`
	// Order date. PST time
	OrderDate string `json:"order_date,omitempty" xml:"order_date,omitempty"`
	// Values: PRODUCT_CUSTOMER_GATHERING,PRODUCT_RETURN_CUSTOMER,PRODUCT_RETURN_WAREHOUSE,PRODUCT_RETURN_SUPPLIER
	LogisticReason string `json:"logistic_reason,omitempty" xml:"logistic_reason,omitempty"`
	// Carrier name
	CarrierName string `json:"carrier_name,omitempty" xml:"carrier_name,omitempty"`
	// The dispute Id
	DisputeId int64 `json:"dispute_id,omitempty" xml:"dispute_id,omitempty"`
}

var poolLogisticOrderCreationForRmaRequest = sync.Pool{
	New: func() any {
		return new(LogisticOrderCreationForRmaRequest)
	},
}

// GetLogisticOrderCreationForRmaRequest() 从对象池中获取LogisticOrderCreationForRmaRequest
func GetLogisticOrderCreationForRmaRequest() *LogisticOrderCreationForRmaRequest {
	return poolLogisticOrderCreationForRmaRequest.Get().(*LogisticOrderCreationForRmaRequest)
}

// ReleaseLogisticOrderCreationForRmaRequest 释放LogisticOrderCreationForRmaRequest
func ReleaseLogisticOrderCreationForRmaRequest(v *LogisticOrderCreationForRmaRequest) {
	v.TrackingCode = ""
	v.OrderDate = ""
	v.LogisticReason = ""
	v.CarrierName = ""
	v.DisputeId = 0
	poolLogisticOrderCreationForRmaRequest.Put(v)
}
