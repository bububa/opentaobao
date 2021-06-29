package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔订购公众号 API请求
taobao.jst.sms.officialaccount.order

聚石塔订购公众号接口
*/
type TaobaoJstSmsOfficialaccountOrderRequest struct {
    model.Params
    // 聚石塔公众号订购
    _orderOfficialAccountRequest   *OrderOfficialAccountRequest
}

// 初始化TaobaoJstSmsOfficialaccountOrderRequest对象
func NewTaobaoJstSmsOfficialaccountOrderRequest() *TaobaoJstSmsOfficialaccountOrderRequest{
    return &TaobaoJstSmsOfficialaccountOrderRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountOrderRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.order"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountOrderRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderOfficialAccountRequest Setter
// 聚石塔公众号订购
func (r *TaobaoJstSmsOfficialaccountOrderRequest) SetOrderOfficialAccountRequest(_orderOfficialAccountRequest *OrderOfficialAccountRequest) error {
    r._orderOfficialAccountRequest = _orderOfficialAccountRequest
    r.Set("order_official_account_request", _orderOfficialAccountRequest)
    return nil
}

// OrderOfficialAccountRequest Getter
func (r TaobaoJstSmsOfficialaccountOrderRequest) GetOrderOfficialAccountRequest() *OrderOfficialAccountRequest {
    return r._orderOfficialAccountRequest
}
