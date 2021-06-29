package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查看订单详情 API请求
taobao.trade.drug.order.get

商家查看订单详情
*/
type TaobaoTradeDrugOrderGetRequest struct {
    model.Params
    // 订单id
    _orderId   int64
}

// 初始化TaobaoTradeDrugOrderGetRequest对象
func NewTaobaoTradeDrugOrderGetRequest() *TaobaoTradeDrugOrderGetRequest{
    return &TaobaoTradeDrugOrderGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTradeDrugOrderGetRequest) GetApiMethodName() string {
    return "taobao.trade.drug.order.get"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTradeDrugOrderGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单id
func (r *TaobaoTradeDrugOrderGetRequest) SetOrderId(_orderId int64) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r TaobaoTradeDrugOrderGetRequest) GetOrderId() int64 {
    return r._orderId
}
