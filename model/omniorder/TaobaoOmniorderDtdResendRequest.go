package omniorder

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
门店自送重发码 APIRequest
taobao.omniorder.dtd.resend

该接口触发对门店自送发码短信进行重发，码内容不变，接受码的手机号也不变。每个码限制每日重发一次，总共重发5次
*/
type TaobaoOmniorderDtdResendRequest struct {
    model.Params

    // 淘宝主订单ID
    mainOrderId   int64 

}

func NewTaobaoOmniorderDtdResendRequest() *TaobaoOmniorderDtdResendRequest{
    return &TaobaoOmniorderDtdResendRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOmniorderDtdResendRequest) GetApiMethodName() string {
    return "taobao.omniorder.dtd.resend"
}

func (r TaobaoOmniorderDtdResendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOmniorderDtdResendRequest) SetMainOrderId(mainOrderId int64) error {
    r.mainOrderId = mainOrderId
    r.Set("main_order_id", mainOrderId)
    return nil
}

func (r TaobaoOmniorderDtdResendRequest) GetMainOrderId() int64 {
    return r.mainOrderId
}

