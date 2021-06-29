package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
判断支付宝用户是否绑定淘宝账号 API请求
alibaba.aliqin.flow.alipay.isbindingtbaccount

判断支付宝用户是否绑定淘宝账号
*/
type AlibabaAliqinFlowAlipayIsbindingtbaccountRequest struct {
    model.Params
    // 支付宝ID
    _alipayId   string
}

// 初始化AlibabaAliqinFlowAlipayIsbindingtbaccountRequest对象
func NewAlibabaAliqinFlowAlipayIsbindingtbaccountRequest() *AlibabaAliqinFlowAlipayIsbindingtbaccountRequest{
    return &AlibabaAliqinFlowAlipayIsbindingtbaccountRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowAlipayIsbindingtbaccountRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.alipay.isbindingtbaccount"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowAlipayIsbindingtbaccountRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayId Setter
// 支付宝ID
func (r *AlibabaAliqinFlowAlipayIsbindingtbaccountRequest) SetAlipayId(_alipayId string) error {
    r._alipayId = _alipayId
    r.Set("alipay_id", _alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaAliqinFlowAlipayIsbindingtbaccountRequest) GetAlipayId() string {
    return r._alipayId
}
