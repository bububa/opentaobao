package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
闲鱼平台商户代扣 APIRequest
alibaba.idle.agreement.pay

闲鱼平台代扣能力：用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
*/
type AlibabaIdleAgreementPayRequest struct {
    model.Params

    // 协议代扣参数
    agreementPayParam   *AgreementPayParam 

}

func NewAlibabaIdleAgreementPayRequest() *AlibabaIdleAgreementPayRequest{
    return &AlibabaIdleAgreementPayRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdleAgreementPayRequest) GetApiMethodName() string {
    return "alibaba.idle.agreement.pay"
}

func (r AlibabaIdleAgreementPayRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdleAgreementPayRequest) SetAgreementPayParam(agreementPayParam *AgreementPayParam) error {
    r.agreementPayParam = agreementPayParam
    r.Set("agreement_pay_param", agreementPayParam)
    return nil
}

func (r AlibabaIdleAgreementPayRequest) GetAgreementPayParam() *AgreementPayParam {
    return r.agreementPayParam
}

