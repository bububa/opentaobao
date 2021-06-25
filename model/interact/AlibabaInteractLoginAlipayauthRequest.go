package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
双11到店互动花呗红包获取token鉴权接口 APIRequest
alibaba.interact.login.alipayauth

双11到店互动花呗红包获取token鉴权接口
*/
type AlibabaInteractLoginAlipayauthRequest struct {
    model.Params

}

func NewAlibabaInteractLoginAlipayauthRequest() *AlibabaInteractLoginAlipayauthRequest{
    return &AlibabaInteractLoginAlipayauthRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractLoginAlipayauthRequest) GetApiMethodName() string {
    return "alibaba.interact.login.alipayauth"
}

func (r AlibabaInteractLoginAlipayauthRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


