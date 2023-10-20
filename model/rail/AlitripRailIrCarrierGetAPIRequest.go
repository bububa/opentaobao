package rail

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripRailIrCarrierGetAPIRequest 国际火车票铁路承运公司查询 API请求
// alitrip.rail.ir.carrier.get
//
// 国际火车票提供给代理商用于查询标准铁路承运公司carrier信息，用于代理商自己的carrier与飞猪平台的carrier做映射
type AlitripRailIrCarrierGetAPIRequest struct {
	model.Params
	// 商家id
	_agentId int64
}

// NewAlitripRailIrCarrierGetRequest 初始化AlitripRailIrCarrierGetAPIRequest对象
func NewAlitripRailIrCarrierGetRequest() *AlitripRailIrCarrierGetAPIRequest {
	return &AlitripRailIrCarrierGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripRailIrCarrierGetAPIRequest) Reset() {
	r._agentId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripRailIrCarrierGetAPIRequest) GetApiMethodName() string {
	return "alitrip.rail.ir.carrier.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripRailIrCarrierGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripRailIrCarrierGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAgentId is AgentId Setter
// 商家id
func (r *AlitripRailIrCarrierGetAPIRequest) SetAgentId(_agentId int64) error {
	r._agentId = _agentId
	r.Set("agent_id", _agentId)
	return nil
}

// GetAgentId AgentId Getter
func (r AlitripRailIrCarrierGetAPIRequest) GetAgentId() int64 {
	return r._agentId
}

var poolAlitripRailIrCarrierGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripRailIrCarrierGetRequest()
	},
}

// GetAlitripRailIrCarrierGetRequest 从 sync.Pool 获取 AlitripRailIrCarrierGetAPIRequest
func GetAlitripRailIrCarrierGetAPIRequest() *AlitripRailIrCarrierGetAPIRequest {
	return poolAlitripRailIrCarrierGetAPIRequest.Get().(*AlitripRailIrCarrierGetAPIRequest)
}

// ReleaseAlitripRailIrCarrierGetAPIRequest 将 AlitripRailIrCarrierGetAPIRequest 放入 sync.Pool
func ReleaseAlitripRailIrCarrierGetAPIRequest(v *AlitripRailIrCarrierGetAPIRequest) {
	v.Reset()
	poolAlitripRailIrCarrierGetAPIRequest.Put(v)
}
