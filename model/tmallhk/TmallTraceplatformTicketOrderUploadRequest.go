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
type TmallTraceplatformTicketOrderUploadRequest struct {
    model.Params
    // 上传小票参数
    ticketOrder   *TicketOrderUpdator
}

// 初始化TmallTraceplatformTicketOrderUploadRequest对象
func NewTmallTraceplatformTicketOrderUploadRequest() *TmallTraceplatformTicketOrderUploadRequest{
    return &TmallTraceplatformTicketOrderUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TmallTraceplatformTicketOrderUploadRequest) GetApiMethodName() string {
    return "tmall.traceplatform.ticket.order.upload"
}

// IRequest interface 方法, 获取API参数
func (r TmallTraceplatformTicketOrderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TicketOrder Setter
// 上传小票参数
func (r *TmallTraceplatformTicketOrderUploadRequest) SetTicketOrder(ticketOrder *TicketOrderUpdator) error {
    r.ticketOrder = ticketOrder
    r.Set("ticket_order", ticketOrder)
    return nil
}

// TicketOrder Getter
func (r TmallTraceplatformTicketOrderUploadRequest) GetTicketOrder() *TicketOrderUpdator {
    return r.ticketOrder
}
