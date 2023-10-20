package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyApproveAPIRequest 销售改签确认 API请求
// alitrip.agent.flight.sell.modify.approve
//
// 销售改签确认
type AlitripAgentFlightSellModifyApproveAPIRequest struct {
	model.Params
	// 入参对象
	_param *ModifyApproveRequestDto
}

// NewAlitripAgentFlightSellModifyApproveRequest 初始化AlitripAgentFlightSellModifyApproveAPIRequest对象
func NewAlitripAgentFlightSellModifyApproveRequest() *AlitripAgentFlightSellModifyApproveAPIRequest {
	return &AlitripAgentFlightSellModifyApproveAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightSellModifyApproveAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyApproveAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.approve"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyApproveAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellModifyApproveAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参对象
func (r *AlitripAgentFlightSellModifyApproveAPIRequest) SetParam(_param *ModifyApproveRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripAgentFlightSellModifyApproveAPIRequest) GetParam() *ModifyApproveRequestDto {
	return r._param
}

var poolAlitripAgentFlightSellModifyApproveAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightSellModifyApproveRequest()
	},
}

// GetAlitripAgentFlightSellModifyApproveRequest 从 sync.Pool 获取 AlitripAgentFlightSellModifyApproveAPIRequest
func GetAlitripAgentFlightSellModifyApproveAPIRequest() *AlitripAgentFlightSellModifyApproveAPIRequest {
	return poolAlitripAgentFlightSellModifyApproveAPIRequest.Get().(*AlitripAgentFlightSellModifyApproveAPIRequest)
}

// ReleaseAlitripAgentFlightSellModifyApproveAPIRequest 将 AlitripAgentFlightSellModifyApproveAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightSellModifyApproveAPIRequest(v *AlitripAgentFlightSellModifyApproveAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightSellModifyApproveAPIRequest.Put(v)
}
