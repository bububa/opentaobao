package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyBackfillAPIRequest 销售改签回填 API请求
// alitrip.agent.flight.sell.modify.backfill
//
// 销售改签回填
type AlitripAgentFlightSellModifyBackfillAPIRequest struct {
	model.Params
	// 入参
	_param *ModifyBackFillRequestDto
}

// NewAlitripAgentFlightSellModifyBackfillRequest 初始化AlitripAgentFlightSellModifyBackfillAPIRequest对象
func NewAlitripAgentFlightSellModifyBackfillRequest() *AlitripAgentFlightSellModifyBackfillAPIRequest {
	return &AlitripAgentFlightSellModifyBackfillAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightSellModifyBackfillAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyBackfillAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.backfill"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyBackfillAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellModifyBackfillAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlitripAgentFlightSellModifyBackfillAPIRequest) SetParam(_param *ModifyBackFillRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripAgentFlightSellModifyBackfillAPIRequest) GetParam() *ModifyBackFillRequestDto {
	return r._param
}

var poolAlitripAgentFlightSellModifyBackfillAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightSellModifyBackfillRequest()
	},
}

// GetAlitripAgentFlightSellModifyBackfillRequest 从 sync.Pool 获取 AlitripAgentFlightSellModifyBackfillAPIRequest
func GetAlitripAgentFlightSellModifyBackfillAPIRequest() *AlitripAgentFlightSellModifyBackfillAPIRequest {
	return poolAlitripAgentFlightSellModifyBackfillAPIRequest.Get().(*AlitripAgentFlightSellModifyBackfillAPIRequest)
}

// ReleaseAlitripAgentFlightSellModifyBackfillAPIRequest 将 AlitripAgentFlightSellModifyBackfillAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightSellModifyBackfillAPIRequest(v *AlitripAgentFlightSellModifyBackfillAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightSellModifyBackfillAPIRequest.Put(v)
}
