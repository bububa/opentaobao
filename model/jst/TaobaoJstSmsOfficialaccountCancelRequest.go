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
type TaobaoJstSmsOfficialaccountCancelRequest struct {
    model.Params
    // 取消公众号订购请求
    cancelOrderRequest   *CancelOrderRequest
}

// 初始化TaobaoJstSmsOfficialaccountCancelRequest对象
func NewTaobaoJstSmsOfficialaccountCancelRequest() *TaobaoJstSmsOfficialaccountCancelRequest{
    return &TaobaoJstSmsOfficialaccountCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoJstSmsOfficialaccountCancelRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.cancel"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoJstSmsOfficialaccountCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CancelOrderRequest Setter
// 取消公众号订购请求
func (r *TaobaoJstSmsOfficialaccountCancelRequest) SetCancelOrderRequest(cancelOrderRequest *CancelOrderRequest) error {
    r.cancelOrderRequest = cancelOrderRequest
    r.Set("cancel_order_request", cancelOrderRequest)
    return nil
}

// CancelOrderRequest Getter
func (r TaobaoJstSmsOfficialaccountCancelRequest) GetCancelOrderRequest() *CancelOrderRequest {
    return r.cancelOrderRequest
}
