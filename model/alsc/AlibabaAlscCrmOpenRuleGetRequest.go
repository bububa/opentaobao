package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询规则 APIRequest
alibaba.alsc.crm.open.rule.get

查询会员规则
*/
type AlibabaAlscCrmOpenRuleGetRequest struct {
    model.Params

    // 入参
    paramRuleOpenReq   *RuleOpenReq 

}

func NewAlibabaAlscCrmOpenRuleGetRequest() *AlibabaAlscCrmOpenRuleGetRequest{
    return &AlibabaAlscCrmOpenRuleGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscCrmOpenRuleGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.rule.get"
}

func (r AlibabaAlscCrmOpenRuleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscCrmOpenRuleGetRequest) SetParamRuleOpenReq(paramRuleOpenReq *RuleOpenReq) error {
    r.paramRuleOpenReq = paramRuleOpenReq
    r.Set("param_rule_open_req", paramRuleOpenReq)
    return nil
}

func (r AlibabaAlscCrmOpenRuleGetRequest) GetParamRuleOpenReq() *RuleOpenReq {
    return r.paramRuleOpenReq
}

