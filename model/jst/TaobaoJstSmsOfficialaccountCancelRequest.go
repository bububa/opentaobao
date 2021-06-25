package jst

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
聚石塔取消公众号订购 APIRequest
taobao.jst.sms.officialaccount.cancel

聚石塔取消公众号订购
*/
type TaobaoJstSmsOfficialaccountCancelRequest struct {
    model.Params

    // 取消公众号订购请求
    cancelOrderRequest   *CancelOrderRequest 

}

func NewTaobaoJstSmsOfficialaccountCancelRequest() *TaobaoJstSmsOfficialaccountCancelRequest{
    return &TaobaoJstSmsOfficialaccountCancelRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoJstSmsOfficialaccountCancelRequest) GetApiMethodName() string {
    return "taobao.jst.sms.officialaccount.cancel"
}

func (r TaobaoJstSmsOfficialaccountCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoJstSmsOfficialaccountCancelRequest) SetCancelOrderRequest(cancelOrderRequest *CancelOrderRequest) error {
    r.cancelOrderRequest = cancelOrderRequest
    r.Set("cancel_order_request", cancelOrderRequest)
    return nil
}

func (r TaobaoJstSmsOfficialaccountCancelRequest) GetCancelOrderRequest() *CancelOrderRequest {
    return r.cancelOrderRequest
}

