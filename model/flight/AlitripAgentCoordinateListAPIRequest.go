package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinatelistAPIRequest 慧飞商家协同单列表查询接口 API请求
// alitrip.agent.coordinate.list
//
// 慧飞商家协同单列表查询接口
type AlitripagentcoordinatelistAPIRequest struct {
	model.Params
	// 协同单列表查询参数
	_coordinationListRequestDto *CoordinationListRequestDto
}

// NewAlitripagentcoordinatelistRequest 初始化AlitripagentcoordinatelistAPIRequest对象
func NewAlitripagentcoordinatelistRequest() *AlitripagentcoordinatelistAPIRequest {
	return &AlitripagentcoordinatelistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentcoordinatelistAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentcoordinatelistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentcoordinatelistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCoordinationListRequestDto is CoordinationListRequestDto Setter
// 协同单列表查询参数
func (r *AlitripagentcoordinatelistAPIRequest) SetCoordinationListRequestDto(_coordinationListRequestDto *CoordinationListRequestDto) error {
	r._coordinationListRequestDto = _coordinationListRequestDto
	r.Set("coordination_list_request_dto", _coordinationListRequestDto)
	return nil
}

// GetCoordinationListRequestDto CoordinationListRequestDto Getter
func (r AlitripagentcoordinatelistAPIRequest) GetCoordinationListRequestDto() *CoordinationListRequestDto {
	return r._coordinationListRequestDto
}
