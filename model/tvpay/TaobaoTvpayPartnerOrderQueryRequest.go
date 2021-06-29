package tvpay

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商户查询订单 API请求
taobao.tvpay.partner.order.query

给商户提供的查询订单状态的API
*/
type TaobaoTvpayPartnerOrderQueryRequest struct {
    model.Params
    // 商户订单号
    _orderNo   string
}

// 初始化TaobaoTvpayPartnerOrderQueryRequest对象
func NewTaobaoTvpayPartnerOrderQueryRequest() *TaobaoTvpayPartnerOrderQueryRequest{
    return &TaobaoTvpayPartnerOrderQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTvpayPartnerOrderQueryRequest) GetApiMethodName() string {
    return "taobao.tvpay.partner.order.query"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTvpayPartnerOrderQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderNo Setter
// 商户订单号
func (r *TaobaoTvpayPartnerOrderQueryRequest) SetOrderNo(_orderNo string) error {
    r._orderNo = _orderNo
    r.Set("order_no", _orderNo)
    return nil
}

// OrderNo Getter
func (r TaobaoTvpayPartnerOrderQueryRequest) GetOrderNo() string {
    return r._orderNo
}
