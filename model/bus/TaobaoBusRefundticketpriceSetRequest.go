package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票退款申请接口 APIRequest
taobao.bus.refundticketprice.set

汽车票代理商利用该接口申请退票
*/
type TaobaoBusRefundticketpriceSetRequest struct {
    model.Params

    // 退票申请入参
    offlineRefundTicketRq   *OfflineRefundTicketPriceRq 

}

func NewTaobaoBusRefundticketpriceSetRequest() *TaobaoBusRefundticketpriceSetRequest{
    return &TaobaoBusRefundticketpriceSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusRefundticketpriceSetRequest) GetApiMethodName() string {
    return "taobao.bus.refundticketprice.set"
}

func (r TaobaoBusRefundticketpriceSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusRefundticketpriceSetRequest) SetOfflineRefundTicketRq(offlineRefundTicketRq *OfflineRefundTicketPriceRq) error {
    r.offlineRefundTicketRq = offlineRefundTicketRq
    r.Set("offline_refund_ticket_rq", offlineRefundTicketRq)
    return nil
}

func (r TaobaoBusRefundticketpriceSetRequest) GetOfflineRefundTicketRq() *OfflineRefundTicketPriceRq {
    return r.offlineRefundTicketRq
}

