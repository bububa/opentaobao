package alscmerchant

import (
	"sync"
)

// ExternalTicketSendRequest 结构体
type ExternalTicketSendRequest struct {
	// 需要发送的码列表，其中code表示串码码值，num表示码的可核销份数
	ExternalTicketCodes []ExternalTicketCode `json:"external_ticket_codes,omitempty" xml:"external_ticket_codes>external_ticket_code,omitempty"`
	// 口碑子订单号, 即alsc_sub_order_no
	OrderNo string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 请求id，唯一标识一次请求
	TicketRequestId string `json:"ticket_request_id,omitempty" xml:"ticket_request_id,omitempty"`
	// 券生效时间,默认为口碑商品配置生效时间
	ValidStart string `json:"valid_start,omitempty" xml:"valid_start,omitempty"`
	// 券过期时间,默认为口碑商品配置失效时间
	ValidEnd string `json:"valid_end,omitempty" xml:"valid_end,omitempty"`
	// 口碑发码通知透传码商，码商需要跟发码通知获取到的参数一致
	SendToken string `json:"send_token,omitempty" xml:"send_token,omitempty"`
	// 口碑商品发货单号
	SendOrderNo string `json:"send_order_no,omitempty" xml:"send_order_no,omitempty"`
	// 三方异步发码是否成功,值为success 代表发码成功,external_ticket_codes 不能为空,fail代表三方发码失败,本地侧凭证发放状态推进到发码失败
	DeliverCodeSuccess string `json:"deliver_code_success,omitempty" xml:"deliver_code_success,omitempty"`
}

var poolExternalTicketSendRequest = sync.Pool{
	New: func() any {
		return new(ExternalTicketSendRequest)
	},
}

// GetExternalTicketSendRequest() 从对象池中获取ExternalTicketSendRequest
func GetExternalTicketSendRequest() *ExternalTicketSendRequest {
	return poolExternalTicketSendRequest.Get().(*ExternalTicketSendRequest)
}

// ReleaseExternalTicketSendRequest 释放ExternalTicketSendRequest
func ReleaseExternalTicketSendRequest(v *ExternalTicketSendRequest) {
	v.ExternalTicketCodes = v.ExternalTicketCodes[:0]
	v.OrderNo = ""
	v.TicketRequestId = ""
	v.ValidStart = ""
	v.ValidEnd = ""
	v.SendToken = ""
	v.SendOrderNo = ""
	v.DeliverCodeSuccess = ""
	poolExternalTicketSendRequest.Put(v)
}
