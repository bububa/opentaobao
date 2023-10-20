package tmallhk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmalltraceplatformticketorderuploadAPIRequest 上传小票数据 API请求
// tmall.traceplatform.ticket.order.upload
//
// upsertOrderBySeller
type TmalltraceplatformticketorderuploadAPIRequest struct {
	model.Params
	// 上传小票参数
	_ticketOrder *TicketOrderUpdator
}

// NewTmalltraceplatformticketorderuploadRequest 初始化TmalltraceplatformticketorderuploadAPIRequest对象
func NewTmalltraceplatformticketorderuploadRequest() *TmalltraceplatformticketorderuploadAPIRequest {
	return &TmalltraceplatformticketorderuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmalltraceplatformticketorderuploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.ticket.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmalltraceplatformticketorderuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmalltraceplatformticketorderuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketOrder is TicketOrder Setter
// 上传小票参数
func (r *TmalltraceplatformticketorderuploadAPIRequest) SetTicketOrder(_ticketOrder *TicketOrderUpdator) error {
	r._ticketOrder = _ticketOrder
	r.Set("ticket_order", _ticketOrder)
	return nil
}

// GetTicketOrder TicketOrder Getter
func (r TmalltraceplatformticketorderuploadAPIRequest) GetTicketOrder() *TicketOrderUpdator {
	return r._ticketOrder
}
