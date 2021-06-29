package ticket

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
【门票API2.0】门票规则信息查询接口 API请求
alitrip.ticket.rule.query

门票规则信息查询接口：返回商家上传的门票规则信息
*/
type AlitripTicketRuleQueryRequest struct {
    model.Params
    // 卖家景点规则编码
    _outRuleId   string
}

// 初始化AlitripTicketRuleQueryRequest对象
func NewAlitripTicketRuleQueryRequest() *AlitripTicketRuleQueryRequest{
    return &AlitripTicketRuleQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTicketRuleQueryRequest) GetApiMethodName() string {
    return "alitrip.ticket.rule.query"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTicketRuleQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OutRuleId Setter
// 卖家景点规则编码
func (r *AlitripTicketRuleQueryRequest) SetOutRuleId(_outRuleId string) error {
    r._outRuleId = _outRuleId
    r.Set("out_rule_id", _outRuleId)
    return nil
}

// OutRuleId Getter
func (r AlitripTicketRuleQueryRequest) GetOutRuleId() string {
    return r._outRuleId
}
