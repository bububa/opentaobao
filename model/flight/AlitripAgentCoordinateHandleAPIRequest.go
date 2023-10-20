package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateHandleAPIRequest 慧飞商家协同单接手接口 API请求
// alitrip.agent.coordinate.handle
//
// 慧飞商家协同单接手接口
type AlitripAgentCoordinateHandleAPIRequest struct {
	model.Params
	// 协同单接手参数
	_handleDto *HandleDto
}

// NewAlitripAgentCoordinateHandleRequest 初始化AlitripAgentCoordinateHandleAPIRequest对象
func NewAlitripAgentCoordinateHandleRequest() *AlitripAgentCoordinateHandleAPIRequest {
	return &AlitripAgentCoordinateHandleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentCoordinateHandleAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.handle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentCoordinateHandleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentCoordinateHandleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHandleDto is HandleDto Setter
// 协同单接手参数
func (r *AlitripAgentCoordinateHandleAPIRequest) SetHandleDto(_handleDto *HandleDto) error {
	r._handleDto = _handleDto
	r.Set("handle_dto", _handleDto)
	return nil
}

// GetHandleDto HandleDto Getter
func (r AlitripAgentCoordinateHandleAPIRequest) GetHandleDto() *HandleDto {
	return r._handleDto
}
