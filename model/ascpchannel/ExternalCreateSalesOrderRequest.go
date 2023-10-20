package ascpchannel

import (
	"sync"
)

// ExternalCreateSalesOrderRequest 结构体
type ExternalCreateSalesOrderRequest struct {
	// 子订单列表
	OutSubOrders []ExternalCreateSubSalesOrderRequest `json:"out_sub_orders,omitempty" xml:"out_sub_orders>external_create_sub_sales_order_request,omitempty"`
	// 币种
	CurrencyType string `json:"currency_type,omitempty" xml:"currency_type,omitempty"`
	// 经销、代销、寄售
	SalesMode string `json:"sales_mode,omitempty" xml:"sales_mode,omitempty"`
	// 外部订单号
	OutOrderNo string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 二级渠道
	SubChannelCode string `json:"sub_channel_code,omitempty" xml:"sub_channel_code,omitempty"`
	// 授权渠道(市场)
	ChannelCode string `json:"channel_code,omitempty" xml:"channel_code,omitempty"`
	// 扩展字段
	Attributes string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// 收货人信息
	Receiver *ExternalReceiverRequest `json:"receiver,omitempty" xml:"receiver,omitempty"`
	// 发货人
	Sender *ExternalSenderRequest `json:"sender,omitempty" xml:"sender,omitempty"`
}

var poolExternalCreateSalesOrderRequest = sync.Pool{
	New: func() any {
		return new(ExternalCreateSalesOrderRequest)
	},
}

// GetExternalCreateSalesOrderRequest() 从对象池中获取ExternalCreateSalesOrderRequest
func GetExternalCreateSalesOrderRequest() *ExternalCreateSalesOrderRequest {
	return poolExternalCreateSalesOrderRequest.Get().(*ExternalCreateSalesOrderRequest)
}

// ReleaseExternalCreateSalesOrderRequest 释放ExternalCreateSalesOrderRequest
func ReleaseExternalCreateSalesOrderRequest(v *ExternalCreateSalesOrderRequest) {
	v.OutSubOrders = v.OutSubOrders[:0]
	v.CurrencyType = ""
	v.SalesMode = ""
	v.OutOrderNo = ""
	v.SubChannelCode = ""
	v.ChannelCode = ""
	v.Attributes = ""
	v.Receiver = nil
	v.Sender = nil
	poolExternalCreateSalesOrderRequest.Put(v)
}
