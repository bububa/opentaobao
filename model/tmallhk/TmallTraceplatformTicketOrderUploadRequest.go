package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小票数据 API请求
tmall.traceplatform.ticket.order.upload

upsertOrderBySeller
*/
type TmallTraceplatformTicketOrderUploadAPIRequest struct {
    model.Params
    // 上传小票参数
    _ticketOrder   *TicketOrderUpdator
}

// 初始化TmallTraceplatformTicketOrderUploadAPIRequest对象
func NewTmallTraceplatformTicketOrderUploadRequest() *TmallTraceplatformTicketOrderUploadAPIRequest{
    return &TmallTraceplatformTicketOrderUploadAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraceplatformTicketOrderUploadAPIRequest) GetApiMethodName() string {
    return "tmall.traceplatform.ticket.order.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraceplatformTicketOrderUploadAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketOrder Setter
// 上传小票参数
func (r *TmallTraceplatformTicketOrderUploadAPIRequest) SetTicketOrder(_ticketOrder *TicketOrderUpdator) error {
    r._ticketOrder = _ticketOrder
    r.Set("ticket_order", _ticketOrder)
    return nil
}

// TicketOrder Getter
func (r TmallTraceplatformTicketOrderUploadAPIRequest) GetTicketOrder() *TicketOrderUpdator {
    return r._ticketOrder
}
