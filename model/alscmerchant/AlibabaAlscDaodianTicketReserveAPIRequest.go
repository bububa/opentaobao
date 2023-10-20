package alscmerchant

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscDaodianTicketReserveAPIRequest 外部券冲正 API请求
// alibaba.alsc.daodian.ticket.reserve
//
// 外部券冲正
type AlibabaAlscDaodianTicketReserveAPIRequest struct {
	model.Params
	// 凭证冲正请求
	_ticketReverseTopRequest *TicketReverseTopRequest
}

// NewAlibabaAlscDaodianTicketReserveRequest 初始化AlibabaAlscDaodianTicketReserveAPIRequest对象
func NewAlibabaAlscDaodianTicketReserveRequest() *AlibabaAlscDaodianTicketReserveAPIRequest {
	return &AlibabaAlscDaodianTicketReserveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlscDaodianTicketReserveAPIRequest) Reset() {
	r._ticketReverseTopRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscDaodianTicketReserveAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.daodian.ticket.reserve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscDaodianTicketReserveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlscDaodianTicketReserveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketReverseTopRequest is TicketReverseTopRequest Setter
// 凭证冲正请求
func (r *AlibabaAlscDaodianTicketReserveAPIRequest) SetTicketReverseTopRequest(_ticketReverseTopRequest *TicketReverseTopRequest) error {
	r._ticketReverseTopRequest = _ticketReverseTopRequest
	r.Set("ticket_reverse_top_request", _ticketReverseTopRequest)
	return nil
}

// GetTicketReverseTopRequest TicketReverseTopRequest Getter
func (r AlibabaAlscDaodianTicketReserveAPIRequest) GetTicketReverseTopRequest() *TicketReverseTopRequest {
	return r._ticketReverseTopRequest
}

var poolAlibabaAlscDaodianTicketReserveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlscDaodianTicketReserveRequest()
	},
}

// GetAlibabaAlscDaodianTicketReserveRequest 从 sync.Pool 获取 AlibabaAlscDaodianTicketReserveAPIRequest
func GetAlibabaAlscDaodianTicketReserveAPIRequest() *AlibabaAlscDaodianTicketReserveAPIRequest {
	return poolAlibabaAlscDaodianTicketReserveAPIRequest.Get().(*AlibabaAlscDaodianTicketReserveAPIRequest)
}

// ReleaseAlibabaAlscDaodianTicketReserveAPIRequest 将 AlibabaAlscDaodianTicketReserveAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlscDaodianTicketReserveAPIRequest(v *AlibabaAlscDaodianTicketReserveAPIRequest) {
	v.Reset()
	poolAlibabaAlscDaodianTicketReserveAPIRequest.Put(v)
}
