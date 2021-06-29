package flight

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
廉航辅营正向订单出货接口 API请求
alitrip.tripvp.agent.order.issue

廉航辅营正向订单出货接口
*/
type AlitripTripvpAgentOrderIssueRequest struct {
    model.Params
    // 代理商ID
    _agentId   int64
    // 出货通知信息
    _issueProOrderVo   *IssueProOrderVo
}

// 初始化AlitripTripvpAgentOrderIssueRequest对象
func NewAlitripTripvpAgentOrderIssueRequest() *AlitripTripvpAgentOrderIssueRequest{
    return &AlitripTripvpAgentOrderIssueRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlitripTripvpAgentOrderIssueRequest) GetApiMethodName() string {
    return "alitrip.tripvp.agent.order.issue"
}

// IRequest interface 方法, 获取API参数
func (r AlitripTripvpAgentOrderIssueRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AgentId Setter
// 代理商ID
func (r *AlitripTripvpAgentOrderIssueRequest) SetAgentId(_agentId int64) error {
    r._agentId = _agentId
    r.Set("agent_id", _agentId)
    return nil
}

// AgentId Getter
func (r AlitripTripvpAgentOrderIssueRequest) GetAgentId() int64 {
    return r._agentId
}
// IssueProOrderVo Setter
// 出货通知信息
func (r *AlitripTripvpAgentOrderIssueRequest) SetIssueProOrderVo(_issueProOrderVo *IssueProOrderVo) error {
    r._issueProOrderVo = _issueProOrderVo
    r.Set("issue_pro_order_vo", _issueProOrderVo)
    return nil
}

// IssueProOrderVo Getter
func (r AlitripTripvpAgentOrderIssueRequest) GetIssueProOrderVo() *IssueProOrderVo {
    return r._issueProOrderVo
}
