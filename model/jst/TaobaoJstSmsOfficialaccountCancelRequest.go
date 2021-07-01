package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔取消公众号订购 API请求
taobao.jst.sms.officialaccount.cancel

聚石塔取消公众号订购
*/
type TaobaoJstSmsOfficialaccountCancelAPIRequest struct {
    model.Params
    // 取消公众号订购请求
    _cancelOrderRequest   *CancelOrderRequest
}

// 初始化TaobaoJstSmsOfficialaccountCancelAPIRequest对象
func NewTaobaoJstSmsOfficialaccountCancelRequest() *TaobaoJstSmsOfficialaccountCancelAPIRequest{
    return &TaobaoJstSmsOfficialaccountCancelAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountCancelAPIRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountCancelAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CancelOrderRequest Setter
// 取消公众号订购请求
func (r *TaobaoJstSmsOfficialaccountCancelAPIRequest) SetCancelOrderRequest(_cancelOrderRequest *CancelOrderRequest) error {
    r._cancelOrderRequest = _cancelOrderRequest
    r.Set("cancel_order_request", _cancelOrderRequest)
    return nil
}

// CancelOrderRequest Getter
func (r TaobaoJstSmsOfficialaccountCancelAPIRequest) GetCancelOrderRequest() *CancelOrderRequest {
    return r._cancelOrderRequest
}
