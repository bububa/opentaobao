package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票退款申请接口 API请求
taobao.bus.refundticketprice.set

汽车票代理商利用该接口申请退票
*/
type TaobaoBusRefundticketpriceSetRequest struct {
    model.Params
    // 退票申请入参
    offlineRefundTicketRq   *OfflineRefundTicketPriceRq
}

// 初始化TaobaoBusRefundticketpriceSetRequest对象
func NewTaobaoBusRefundticketpriceSetRequest() *TaobaoBusRefundticketpriceSetRequest{
    return &TaobaoBusRefundticketpriceSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusRefundticketpriceSetRequest) GetApiMethodName() string {
    return "taobao.bus.refundticketprice.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusRefundticketpriceSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OfflineRefundTicketRq Setter
// 退票申请入参
func (r *TaobaoBusRefundticketpriceSetRequest) SetOfflineRefundTicketRq(offlineRefundTicketRq *OfflineRefundTicketPriceRq) error {
    r.offlineRefundTicketRq = offlineRefundTicketRq
    r.Set("offline_refund_ticket_rq", offlineRefundTicketRq)
    return nil
}

// OfflineRefundTicketRq Getter
func (r TaobaoBusRefundticketpriceSetRequest) GetOfflineRefundTicketRq() *OfflineRefundTicketPriceRq {
    return r.offlineRefundTicketRq
}
