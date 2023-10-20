package tmallhk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TmallTraceplatformTicketOrderUploadAPIRequest 上传小票数据 API请求
// tmall.traceplatform.ticket.order.upload
//
// upsertOrderBySeller
type TmallTraceplatformTicketOrderUploadAPIRequest struct {
	model.Params
	// 上传小票参数
	_ticketOrder *TicketOrderUpdator
}

// NewTmallTraceplatformTicketOrderUploadRequest 初始化TmallTraceplatformTicketOrderUploadAPIRequest对象
func NewTmallTraceplatformTicketOrderUploadRequest() *TmallTraceplatformTicketOrderUploadAPIRequest {
	return &TmallTraceplatformTicketOrderUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TmallTraceplatformTicketOrderUploadAPIRequest) Reset() {
	r._ticketOrder = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallTraceplatformTicketOrderUploadAPIRequest) GetApiMethodName() string {
	return "tmall.traceplatform.ticket.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallTraceplatformTicketOrderUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallTraceplatformTicketOrderUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTicketOrder is TicketOrder Setter
// 上传小票参数
func (r *TmallTraceplatformTicketOrderUploadAPIRequest) SetTicketOrder(_ticketOrder *TicketOrderUpdator) error {
	r._ticketOrder = _ticketOrder
	r.Set("ticket_order", _ticketOrder)
	return nil
}

// GetTicketOrder TicketOrder Getter
func (r TmallTraceplatformTicketOrderUploadAPIRequest) GetTicketOrder() *TicketOrderUpdator {
	return r._ticketOrder
}

var poolTmallTraceplatformTicketOrderUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewTmallTraceplatformTicketOrderUploadRequest()
	},
}

// GetTmallTraceplatformTicketOrderUploadRequest 从 sync.Pool 获取 TmallTraceplatformTicketOrderUploadAPIRequest
func GetTmallTraceplatformTicketOrderUploadAPIRequest() *TmallTraceplatformTicketOrderUploadAPIRequest {
	return poolTmallTraceplatformTicketOrderUploadAPIRequest.Get().(*TmallTraceplatformTicketOrderUploadAPIRequest)
}

// ReleaseTmallTraceplatformTicketOrderUploadAPIRequest 将 TmallTraceplatformTicketOrderUploadAPIRequest 放入 sync.Pool
func ReleaseTmallTraceplatformTicketOrderUploadAPIRequest(v *TmallTraceplatformTicketOrderUploadAPIRequest) {
	v.Reset()
	poolTmallTraceplatformTicketOrderUploadAPIRequest.Put(v)
}
