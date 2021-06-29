package idle

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
创建代扣计划 APIRequest
alibaba.idle.pay.plan.create

闲鱼平台代扣能力：
1、用户和闲鱼签约代扣协议 服务商通过直付通产品挂载成为闲鱼二级商户 来完成用户和服务商的结算
2、创建代扣计划
*/
type AlibabaIdlePayPlanCreateRequest struct {
    model.Params

    // 业务入参
    agreementPayPlanParam   *AgreementPayPlanParam 

}

func NewAlibabaIdlePayPlanCreateRequest() *AlibabaIdlePayPlanCreateRequest{
    return &AlibabaIdlePayPlanCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIdlePayPlanCreateRequest) GetApiMethodName() string {
    return "alibaba.idle.pay.plan.create"
}

func (r AlibabaIdlePayPlanCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaIdlePayPlanCreateRequest) SetAgreementPayPlanParam(agreementPayPlanParam *AgreementPayPlanParam) error {
    r.agreementPayPlanParam = agreementPayPlanParam
    r.Set("agreement_pay_plan_param", agreementPayPlanParam)
    return nil
}

func (r AlibabaIdlePayPlanCreateRequest) GetAgreementPayPlanParam() *AgreementPayPlanParam {
    return r.agreementPayPlanParam
}

