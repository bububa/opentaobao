package trade

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断该订单能执行的逆向操作集合列表 API请求
cainiao.refund.refundactions.get

判断该订单能执行的逆向操作集合列表
*/
type CainiaoRefundRefundactionsGetRequest struct {
    model.Params
    // 子订单ID
    _orderId   string
}

// 初始化CainiaoRefundRefundactionsGetRequest对象
func NewCainiaoRefundRefundactionsGetRequest() *CainiaoRefundRefundactionsGetRequest{
    return &CainiaoRefundRefundactionsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoRefundRefundactionsGetRequest) GetApiMethodName() string {
    return "cainiao.refund.refundactions.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoRefundRefundactionsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 子订单ID
func (r *CainiaoRefundRefundactionsGetRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r CainiaoRefundRefundactionsGetRequest) GetOrderId() string {
    return r._orderId
}
