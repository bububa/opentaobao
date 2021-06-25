package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票规则信息查询接口 APIRequest
alitrip.ticket.rule.query

门票规则信息查询接口：返回商家上传的门票规则信息
*/
type AlitripTicketRuleQueryRequest struct {
    model.Params

    // 卖家景点规则编码
    outRuleId   string 

}

func NewAlitripTicketRuleQueryRequest() *AlitripTicketRuleQueryRequest{
    return &AlitripTicketRuleQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlitripTicketRuleQueryRequest) GetApiMethodName() string {
    return "alitrip.ticket.rule.query"
}

func (r AlitripTicketRuleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlitripTicketRuleQueryRequest) SetOutRuleId(outRuleId string) error {
    r.outRuleId = outRuleId
    r.Set("out_rule_id", outRuleId)
    return nil
}

func (r AlitripTicketRuleQueryRequest) GetOutRuleId() string {
    return r.outRuleId
}

