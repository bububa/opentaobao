package interact

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
双11到店互动花呗红包获取token鉴权接口 API请求
alibaba.interact.login.alipayauth

双11到店互动花呗红包获取token鉴权接口
*/
type AlibabaInteractLoginAlipayauthRequest struct {
    model.Params
}

// 初始化AlibabaInteractLoginAlipayauthRequest对象
func NewAlibabaInteractLoginAlipayauthRequest() *AlibabaInteractLoginAlipayauthRequest{
    return &AlibabaInteractLoginAlipayauthRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractLoginAlipayauthRequest) GetApiMethodName() string {
    return "alibaba.interact.login.alipayauth"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractLoginAlipayauthRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
