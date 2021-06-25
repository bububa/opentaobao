package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值规则下行 APIRequest
alibaba.alsc.crm.recharge.qryrule

储值规则下行
*/
type AlibabaAlscCrmRechargeQryruleRequest struct {
    model.Params

    // 请求对象
    paramPullRechargeRuleByShopReq   *PullRechargeRuleByShopReq 

}

func NewAlibabaAlscCrmRechargeQryruleRequest() *AlibabaAlscCrmRechargeQryruleRequest{
    return &AlibabaAlscCrmRechargeQryruleRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmRechargeQryruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.qryrule"
}

func (r AlibabaAlscCrmRechargeQryruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmRechargeQryruleRequest) SetParamPullRechargeRuleByShopReq(paramPullRechargeRuleByShopReq *PullRechargeRuleByShopReq) error {
    r.paramPullRechargeRuleByShopReq = paramPullRechargeRuleByShopReq
    r.Set("param_pull_recharge_rule_by_shop_req", paramPullRechargeRuleByShopReq)
    return nil
}

func (r AlibabaAlscCrmRechargeQryruleRequest) GetParamPullRechargeRuleByShopReq() *PullRechargeRuleByShopReq {
    return r.paramPullRechargeRuleByShopReq
}

