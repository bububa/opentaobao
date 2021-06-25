package tmallhk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上传小票数据 APIRequest
tmall.traceplatform.ticket.order.upload

upsertOrderBySeller
*/
type TmallTraceplatformTicketOrderUploadRequest struct {
    model.Params

    // 上传小票参数
    ticketOrder   *TicketOrderUpdator 

}

func NewTmallTraceplatformTicketOrderUploadRequest() *TmallTraceplatformTicketOrderUploadRequest{
    return &TmallTraceplatformTicketOrderUploadRequest{
        Params: model.NewParams(),
    }
}

func (r TmallTraceplatformTicketOrderUploadRequest) GetApiMethodName() string {
    return "tmall.traceplatform.ticket.order.upload"
}

func (r TmallTraceplatformTicketOrderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TmallTraceplatformTicketOrderUploadRequest) SetTicketOrder(ticketOrder *TicketOrderUpdator) error {
    r.ticketOrder = ticketOrder
    r.Set("ticket_order", ticketOrder)
    return nil
}

func (r TmallTraceplatformTicketOrderUploadRequest) GetTicketOrder() *TicketOrderUpdator {
    return r.ticketOrder
}

