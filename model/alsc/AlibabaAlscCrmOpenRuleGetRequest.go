package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询规则 API请求
alibaba.alsc.crm.open.rule.get

查询会员规则
*/
type AlibabaAlscCrmOpenRuleGetRequest struct {
    model.Params
    // 入参
    _paramRuleOpenReq   *RuleOpenReq
}

// 初始化AlibabaAlscCrmOpenRuleGetRequest对象
func NewAlibabaAlscCrmOpenRuleGetRequest() *AlibabaAlscCrmOpenRuleGetRequest{
    return &AlibabaAlscCrmOpenRuleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmOpenRuleGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.open.rule.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmOpenRuleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamRuleOpenReq Setter
// 入参
func (r *AlibabaAlscCrmOpenRuleGetRequest) SetParamRuleOpenReq(_paramRuleOpenReq *RuleOpenReq) error {
    r._paramRuleOpenReq = _paramRuleOpenReq
    r.Set("param_rule_open_req", _paramRuleOpenReq)
    return nil
}

// ParamRuleOpenReq Getter
func (r AlibabaAlscCrmOpenRuleGetRequest) GetParamRuleOpenReq() *RuleOpenReq {
    return r._paramRuleOpenReq
}
