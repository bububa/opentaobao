package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
生态pos购后发放权益 APIRequest
alibaba.wdk.pos.afterbuy.benefit.send

生态pos购后发放权益接口开放
*/
type AlibabaWdkPosAfterbuyBenefitSendRequest struct {
    model.Params

    // 入参
    sendBenefitParam   *IsvSendBenefitParam 

}

func NewAlibabaWdkPosAfterbuyBenefitSendRequest() *AlibabaWdkPosAfterbuyBenefitSendRequest{
    return &AlibabaWdkPosAfterbuyBenefitSendRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkPosAfterbuyBenefitSendRequest) GetApiMethodName() string {
    return "alibaba.wdk.pos.afterbuy.benefit.send"
}

func (r AlibabaWdkPosAfterbuyBenefitSendRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkPosAfterbuyBenefitSendRequest) SetSendBenefitParam(sendBenefitParam *IsvSendBenefitParam) error {
    r.sendBenefitParam = sendBenefitParam
    r.Set("send_benefit_param", sendBenefitParam)
    return nil
}

func (r AlibabaWdkPosAfterbuyBenefitSendRequest) GetSendBenefitParam() *IsvSendBenefitParam {
    return r.sendBenefitParam
}

