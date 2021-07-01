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
type CainiaoRefundRefundactionsGetAPIRequest struct {
    model.Params
    // 子订单ID
    _orderId   string
}

// 初始化CainiaoRefundRefundactionsGetAPIRequest对象
func NewCainiaoRefundRefundactionsGetRequest() *CainiaoRefundRefundactionsGetAPIRequest{
    return &CainiaoRefundRefundactionsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoRefundRefundactionsGetAPIRequest) GetApiMethodName() string {
    return "cainiao.refund.refundactions.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoRefundRefundactionsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 子订单ID
func (r *CainiaoRefundRefundactionsGetAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r CainiaoRefundRefundactionsGetAPIRequest) GetOrderId() string {
    return r._orderId
}
