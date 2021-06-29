package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
储值规则下行 API请求
alibaba.alsc.crm.recharge.qryrule

储值规则下行
*/
type AlibabaAlscCrmRechargeQryruleRequest struct {
    model.Params
    // 请求对象
    paramPullRechargeRuleByShopReq   *PullRechargeRuleByShopReq
}

// 初始化AlibabaAlscCrmRechargeQryruleRequest对象
func NewAlibabaAlscCrmRechargeQryruleRequest() *AlibabaAlscCrmRechargeQryruleRequest{
    return &AlibabaAlscCrmRechargeQryruleRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmRechargeQryruleRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.recharge.qryrule"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmRechargeQryruleRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamPullRechargeRuleByShopReq Setter
// 请求对象
func (r *AlibabaAlscCrmRechargeQryruleRequest) SetParamPullRechargeRuleByShopReq(paramPullRechargeRuleByShopReq *PullRechargeRuleByShopReq) error {
    r.paramPullRechargeRuleByShopReq = paramPullRechargeRuleByShopReq
    r.Set("param_pull_recharge_rule_by_shop_req", paramPullRechargeRuleByShopReq)
    return nil
}

// ParamPullRechargeRuleByShopReq Getter
func (r AlibabaAlscCrmRechargeQryruleRequest) GetParamPullRechargeRuleByShopReq() *PullRechargeRuleByShopReq {
    return r.paramPullRechargeRuleByShopReq
}
