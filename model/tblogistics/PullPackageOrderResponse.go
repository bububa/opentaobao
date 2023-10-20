package tblogistics

import (
	"sync"
)

// PullPackageOrderResponse 结构体
type PullPackageOrderResponse struct {
	// 包裹出库单号
	DeliveryOrderCode string `json:"delivery_order_code,omitempty" xml:"delivery_order_code,omitempty"`
	// 物流服务商ID
	MailNo string `json:"mail_no,omitempty" xml:"mail_no,omitempty"`
	// 主订单id，正向出库时下发
	OaidOrderSourceCode string `json:"oaid_order_source_code,omitempty" xml:"oaid_order_source_code,omitempty"`
	// 拓展字段
	ExtendProps string `json:"extend_props,omitempty" xml:"extend_props,omitempty"`
	// 配送公司
	TmsCpCode string `json:"tms_cp_code,omitempty" xml:"tms_cp_code,omitempty"`
	// 包裹入库单号
	EntryOrderCode string `json:"entry_order_code,omitempty" xml:"entry_order_code,omitempty"`
	// 物流主体，例如：TaoTian(淘天)
	LogisticsOwner string `json:"logistics_owner,omitempty" xml:"logistics_owner,omitempty"`
	// 收件人信息，逆向出库时下发
	ReceiverInfo *ContactInfo `json:"receiver_info,omitempty" xml:"receiver_info,omitempty"`
}

var poolPullPackageOrderResponse = sync.Pool{
	New: func() any {
		return new(PullPackageOrderResponse)
	},
}

// GetPullPackageOrderResponse() 从对象池中获取PullPackageOrderResponse
func GetPullPackageOrderResponse() *PullPackageOrderResponse {
	return poolPullPackageOrderResponse.Get().(*PullPackageOrderResponse)
}

// ReleasePullPackageOrderResponse 释放PullPackageOrderResponse
func ReleasePullPackageOrderResponse(v *PullPackageOrderResponse) {
	v.DeliveryOrderCode = ""
	v.MailNo = ""
	v.OaidOrderSourceCode = ""
	v.ExtendProps = ""
	v.TmsCpCode = ""
	v.EntryOrderCode = ""
	v.LogisticsOwner = ""
	v.ReceiverInfo = nil
	poolPullPackageOrderResponse.Put(v)
}
