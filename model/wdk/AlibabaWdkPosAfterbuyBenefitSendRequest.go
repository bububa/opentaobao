package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生态pos购后发放权益 API请求
alibaba.wdk.pos.afterbuy.benefit.send

生态pos购后发放权益接口开放
*/
type AlibabaWdkPosAfterbuyBenefitSendRequest struct {
    model.Params
    // 入参
    _sendBenefitParam   *IsvSendBenefitParam
}

// 初始化AlibabaWdkPosAfterbuyBenefitSendRequest对象
func NewAlibabaWdkPosAfterbuyBenefitSendRequest() *AlibabaWdkPosAfterbuyBenefitSendRequest{
    return &AlibabaWdkPosAfterbuyBenefitSendRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaWdkPosAfterbuyBenefitSendRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.afterbuy.benefit.send"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaWdkPosAfterbuyBenefitSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// SendBenefitParam Setter
// 入参
func (r *AlibabaWdkPosAfterbuyBenefitSendRequest) SetSendBenefitParam(_sendBenefitParam *IsvSendBenefitParam) error {
    r._sendBenefitParam = _sendBenefitParam
    r.Set("send_benefit_param", _sendBenefitParam)
    return nil
}

// SendBenefitParam Getter
func (r AlibabaWdkPosAfterbuyBenefitSendRequest) GetSendBenefitParam() *IsvSendBenefitParam {
    return r._sendBenefitParam
}
