package rail

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripRailIrServiceGetAPIRequest) Reset() {
	r._bizType = 0
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailIrServiceGetAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.ir.service.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailIrServiceGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripRailIrServiceGetAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlitripRailIrServiceGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripRailIrServiceGetRequest()
	},
}

// GetAlitripRailIrServiceGetRequest 从 sync.Pool 获取 AlitripRailIrServiceGetAPIRequest
func GetAlitripRailIrServiceGetAPIRequest() *AlitripRailIrServiceGetAPIRequest {
	return poolAlitripRailIrServiceGetAPIRequest.Get().(*AlitripRailIrServiceGetAPIRequest)
}

// ReleaseAlitripRailIrServiceGetAPIRequest 将 AlitripRailIrServiceGetAPIRequest 放入 sync.Pool
func ReleaseAlitripRailIrServiceGetAPIRequest(v *AlitripRailIrServiceGetAPIRequest) {
	v.Reset()
	poolAlitripRailIrServiceGetAPIRequest.Put(v)
}
