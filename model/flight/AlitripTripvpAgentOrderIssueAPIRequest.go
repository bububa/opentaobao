package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTripvpAgentOrderIssueAPIRequest 廉航辅营正向订单出货接口 API请求
// alitrip.tripvp.agent.order.issue
//
// 廉航辅营正向订单出货接口
type AlitripTripvpAgentOrderIssueAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 出货通知信息
	_issueProOrderVo *IssueProOrderVo
}

// NewAlitripTripvpAgentOrderIssueRequest 初始化AlitripTripvpAgentOrderIssueAPIRequest对象
func NewAlitripTripvpAgentOrderIssueRequest() *AlitripTripvpAgentOrderIssueAPIRequest {
	return &AlitripTripvpAgentOrderIssueAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTripvpAgentOrderIssueAPIRequest) Reset() {
	r._agentId = 0
	r._issueProOrderVo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetApiMethodName() string {
	return "alitrip.tripvp.agent.order.issue"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 代理商ID
func (r *AlitripTripvpAgentOrderIssueAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetAgentId() int64 {
	return r._agentId
}

// SetIssueProOrderVo is IssueProOrderVo Setter
// 出货通知信息
func (r *AlitripTripvpAgentOrderIssueAPIRequest) SetIssueProOrderVo(_issueProOrderVo *IssueProOrderVo) error {
	r._issueProOrderVo = _issueProOrderVo
	r.Set("issue_pro_order_vo", _issueProOrderVo)
	return nil
}

// GetIssueProOrderVo IssueProOrderVo Getter
func (r AlitripTripvpAgentOrderIssueAPIRequest) GetIssueProOrderVo() *IssueProOrderVo {
	return r._issueProOrderVo
}

var poolAlitripTripvpAgentOrderIssueAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTripvpAgentOrderIssueRequest()
	},
}

// GetAlitripTripvpAgentOrderIssueRequest 从 sync.Pool 获取 AlitripTripvpAgentOrderIssueAPIRequest
func GetAlitripTripvpAgentOrderIssueAPIRequest() *AlitripTripvpAgentOrderIssueAPIRequest {
	return poolAlitripTripvpAgentOrderIssueAPIRequest.Get().(*AlitripTripvpAgentOrderIssueAPIRequest)
}

// ReleaseAlitripTripvpAgentOrderIssueAPIRequest 将 AlitripTripvpAgentOrderIssueAPIRequest 放入 sync.Pool
func ReleaseAlitripTripvpAgentOrderIssueAPIRequest(v *AlitripTripvpAgentOrderIssueAPIRequest) {
	v.Reset()
	poolAlitripTripvpAgentOrderIssueAPIRequest.Put(v)
}
