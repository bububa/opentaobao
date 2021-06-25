package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断支付宝用户是否绑定淘宝账号 APIRequest
alibaba.aliqin.flow.alipay.isbindingtbaccount

判断支付宝用户是否绑定淘宝账号
*/
type AlibabaAliqinFlowAlipayIsbindingtbaccountRequest struct {
    model.Params

    // 支付宝ID
    alipayId   string 

}

func NewAlibabaAliqinFlowAlipayIsbindingtbaccountRequest() *AlibabaAliqinFlowAlipayIsbindingtbaccountRequest{
    return &AlibabaAliqinFlowAlipayIsbindingtbaccountRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowAlipayIsbindingtbaccountRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.alipay.isbindingtbaccount"
}

func (r AlibabaAliqinFlowAlipayIsbindingtbaccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowAlipayIsbindingtbaccountRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

func (r AlibabaAliqinFlowAlipayIsbindingtbaccountRequest) GetAlipayId() string {
    return r.alipayId
}

