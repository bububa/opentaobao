package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitriprailircarriergetAPIRequest 国际火车票铁路承运公司查询 API请求
// alitrip.rail.ir.carrier.get
//
// 国际火车票提供给代理商用于查询标准铁路承运公司carrier信息，用于代理商自己的carrier与飞猪平台的carrier做映射
type AlitriprailircarriergetAPIRequest struct {
	model.Params
	// 商家id
	_agentId int64
}

// NewAlitriprailircarriergetRequest 初始化AlitriprailircarriergetAPIRequest对象
func NewAlitriprailircarriergetRequest() *AlitriprailircarriergetAPIRequest {
	return &AlitriprailircarriergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitriprailircarriergetAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.ir.carrier.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitriprailircarriergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitriprailircarriergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 商家id
func (r *AlitriprailircarriergetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitriprailircarriergetAPIRequest) GetAgentId() int64 {
	return r._agentId
}
