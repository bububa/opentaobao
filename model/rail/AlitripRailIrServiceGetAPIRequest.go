package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailIrServiceGetAPIRequest 国际火车票仓位坐席查询 API请求
// alitrip.rail.ir.service.get
//
// 国际火车票标准仓位坐席查询
type AlitripRailIrServiceGetAPIRequest struct {
	model.Params
	// 6代表境外火车票
	_bizType int64
	// 代理商id
	_agentId int64
}

// NewAlitripRailIrServiceGetRequest 初始化AlitripRailIrServiceGetAPIRequest对象
func NewAlitripRailIrServiceGetRequest() *AlitripRailIrServiceGetAPIRequest {
	return &AlitripRailIrServiceGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailIrServiceGetAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.ir.service.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailIrServiceGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBizType is BizType Setter
// 6代表境外火车票
func (r *AlitripRailIrServiceGetAPIRequest) SetBizType(_bizType int64) error {
	r._bizType = _bizType
	r.Set("biz_type", _bizType)
	return nil
}

// GetBizType BizType Getter
func (r AlitripRailIrServiceGetAPIRequest) GetBizType() int64 {
	return r._bizType
}

// SetAgentId is AgentId Setter
// 代理商id
func (r *AlitripRailIrServiceGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitripRailIrServiceGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
