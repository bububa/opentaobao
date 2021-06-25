package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断该订单能执行的逆向操作集合列表 APIRequest
cainiao.refund.refundactions.get

判断该订单能执行的逆向操作集合列表
*/
type CainiaoRefundRefundactionsGetRequest struct {
    model.Params

    // 子订单ID
    orderId   string 

}

func NewCainiaoRefundRefundactionsGetRequest() *CainiaoRefundRefundactionsGetRequest{
    return &CainiaoRefundRefundactionsGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoRefundRefundactionsGetRequest) GetApiMethodName() string {
    return "cainiao.refund.refundactions.get"
}

func (r CainiaoRefundRefundactionsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoRefundRefundactionsGetRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r CainiaoRefundRefundactionsGetRequest) GetOrderId() string {
    return r.orderId
}

