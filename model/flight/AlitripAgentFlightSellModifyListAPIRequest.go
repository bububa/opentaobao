package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightSellModifyListAPIRequest 销售改签单列表 API请求
// alitrip.agent.flight.sell.modify.list
//
// 销售改签单列表
type AlitripAgentFlightSellModifyListAPIRequest struct {
	model.Params
	// 入参
	_param *ModifyListRequestDto
}

// NewAlitripAgentFlightSellModifyListRequest 初始化AlitripAgentFlightSellModifyListAPIRequest对象
func NewAlitripAgentFlightSellModifyListRequest() *AlitripAgentFlightSellModifyListAPIRequest {
	return &AlitripAgentFlightSellModifyListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightSellModifyListAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyListAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightSellModifyListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 入参
func (r *AlitripAgentFlightSellModifyListAPIRequest) SetParam(_param *ModifyListRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlitripAgentFlightSellModifyListAPIRequest) GetParam() *ModifyListRequestDto {
	return r._param
}

var poolAlitripAgentFlightSellModifyListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightSellModifyListRequest()
	},
}

// GetAlitripAgentFlightSellModifyListRequest 从 sync.Pool 获取 AlitripAgentFlightSellModifyListAPIRequest
func GetAlitripAgentFlightSellModifyListAPIRequest() *AlitripAgentFlightSellModifyListAPIRequest {
	return poolAlitripAgentFlightSellModifyListAPIRequest.Get().(*AlitripAgentFlightSellModifyListAPIRequest)
}

// ReleaseAlitripAgentFlightSellModifyListAPIRequest 将 AlitripAgentFlightSellModifyListAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightSellModifyListAPIRequest(v *AlitripAgentFlightSellModifyListAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightSellModifyListAPIRequest.Put(v)
}
