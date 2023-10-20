package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinatehandleAPIRequest 慧飞商家协同单接手接口 API请求
// alitrip.agent.coordinate.handle
//
// 慧飞商家协同单接手接口
type AlitripagentcoordinatehandleAPIRequest struct {
	model.Params
	// 协同单接手参数
	_handleDto *HandleDto
}

// NewAlitripagentcoordinatehandleRequest 初始化AlitripagentcoordinatehandleAPIRequest对象
func NewAlitripagentcoordinatehandleRequest() *AlitripagentcoordinatehandleAPIRequest {
	return &AlitripagentcoordinatehandleAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentcoordinatehandleAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.handle"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentcoordinatehandleAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentcoordinatehandleAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetHandleDto is HandleDto Setter
// 协同单接手参数
func (r *AlitripagentcoordinatehandleAPIRequest) SetHandleDto(_handleDto *HandleDto) error {
	r._handleDto = _handleDto
	r.Set("handle_dto", _handleDto)
	return nil
}

// GetHandleDto HandleDto Getter
func (r AlitripagentcoordinatehandleAPIRequest) GetHandleDto() *HandleDto {
	return r._handleDto
}
