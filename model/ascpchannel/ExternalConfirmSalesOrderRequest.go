package ascpchannel

import (
	"sync"
)

// ExternalConfirmSalesOrderRequest 结构体
type ExternalConfirmSalesOrderRequest struct {
	// 渠道主单号
	SaleOrderNo string `json:"sale_order_no,omitempty" xml:"sale_order_no,omitempty"`
}

var poolExternalConfirmSalesOrderRequest = sync.Pool{
	New: func() any {
		return new(ExternalConfirmSalesOrderRequest)
	},
}

// GetExternalConfirmSalesOrderRequest() 从对象池中获取ExternalConfirmSalesOrderRequest
func GetExternalConfirmSalesOrderRequest() *ExternalConfirmSalesOrderRequest {
	return poolExternalConfirmSalesOrderRequest.Get().(*ExternalConfirmSalesOrderRequest)
}

// ReleaseExternalConfirmSalesOrderRequest 释放ExternalConfirmSalesOrderRequest
func ReleaseExternalConfirmSalesOrderRequest(v *ExternalConfirmSalesOrderRequest) {
	v.SaleOrderNo = ""
	poolExternalConfirmSalesOrderRequest.Put(v)
}
