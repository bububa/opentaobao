package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询积分规则 API请求
alibaba.alsc.crm.point.rule.get

新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口
*/
type AlibabaAlscCrmPointRuleGetRequest struct {
    model.Params
    // 入参
    paramQueryPointRuleOpenReq   *QueryPointRuleOpenReq
}

// 初始化AlibabaAlscCrmPointRuleGetRequest对象
func NewAlibabaAlscCrmPointRuleGetRequest() *AlibabaAlscCrmPointRuleGetRequest{
    return &AlibabaAlscCrmPointRuleGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscCrmPointRuleGetRequest) GetApiMethodName() string {
    return "alibaba.alsc.crm.point.rule.get"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscCrmPointRuleGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamQueryPointRuleOpenReq Setter
// 入参
func (r *AlibabaAlscCrmPointRuleGetRequest) SetParamQueryPointRuleOpenReq(paramQueryPointRuleOpenReq *QueryPointRuleOpenReq) error {
    r.paramQueryPointRuleOpenReq = paramQueryPointRuleOpenReq
    r.Set("param_query_point_rule_open_req", paramQueryPointRuleOpenReq)
    return nil
}

// ParamQueryPointRuleOpenReq Getter
func (r AlibabaAlscCrmPointRuleGetRequest) GetParamQueryPointRuleOpenReq() *QueryPointRuleOpenReq {
    return r.paramQueryPointRuleOpenReq
}
