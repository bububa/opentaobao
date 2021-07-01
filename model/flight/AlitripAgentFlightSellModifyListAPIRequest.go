package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripAgentFlightSellModifyListAPIRequest
销售改签单列表 API请求
alitrip.agent.flight.sell.modify.list

销售改签单列表 */
type AlitripAgentFlightSellModifyListAPIRequest struct {
	model.Params
	// 入参
	_param *ModifyListRequestDto
}

// NewAlitripAgentFlightSellModifyListRequest 初始化AlitripAgentFlightSellModifyListAPIRequest对象
func NewAlitripAgentFlightSellModifyListRequest() *AlitripAgentFlightSellModifyListAPIRequest {
	return &AlitripAgentFlightSellModifyListAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightSellModifyListAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.sell.modify.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightSellModifyListAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Param Setter
// 入参
func (r *AlitripAgentFlightSellModifyListAPIRequest) SetParam(_param *ModifyListRequestDto) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// Get Param Getter
func (r AlitripAgentFlightSellModifyListAPIRequest) GetParam() *ModifyListRequestDto {
	return r._param
}
