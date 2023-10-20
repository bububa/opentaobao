package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateListAPIRequest 慧飞商家协同单列表查询接口 API请求
// alitrip.agent.coordinate.list
//
// 慧飞商家协同单列表查询接口
type AlitripAgentCoordinateListAPIRequest struct {
	model.Params
	// 协同单列表查询参数
	_coordinationListRequestDto *CoordinationListRequestDto
}

// NewAlitripAgentCoordinateListRequest 初始化AlitripAgentCoordinateListAPIRequest对象
func NewAlitripAgentCoordinateListRequest() *AlitripAgentCoordinateListAPIRequest {
	return &AlitripAgentCoordinateListAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentCoordinateListAPIRequest) Reset() {
	r._coordinationListRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentCoordinateListAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.list"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentCoordinateListAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentCoordinateListAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCoordinationListRequestDto is CoordinationListRequestDto Setter
// 协同单列表查询参数
func (r *AlitripAgentCoordinateListAPIRequest) SetCoordinationListRequestDto(_coordinationListRequestDto *CoordinationListRequestDto) error {
	r._coordinationListRequestDto = _coordinationListRequestDto
	r.Set("coordination_list_request_dto", _coordinationListRequestDto)
	return nil
}

// GetCoordinationListRequestDto CoordinationListRequestDto Getter
func (r AlitripAgentCoordinateListAPIRequest) GetCoordinationListRequestDto() *CoordinationListRequestDto {
	return r._coordinationListRequestDto
}

var poolAlitripAgentCoordinateListAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentCoordinateListRequest()
	},
}

// GetAlitripAgentCoordinateListRequest 从 sync.Pool 获取 AlitripAgentCoordinateListAPIRequest
func GetAlitripAgentCoordinateListAPIRequest() *AlitripAgentCoordinateListAPIRequest {
	return poolAlitripAgentCoordinateListAPIRequest.Get().(*AlitripAgentCoordinateListAPIRequest)
}

// ReleaseAlitripAgentCoordinateListAPIRequest 将 AlitripAgentCoordinateListAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentCoordinateListAPIRequest(v *AlitripAgentCoordinateListAPIRequest) {
	v.Reset()
	poolAlitripAgentCoordinateListAPIRequest.Put(v)
}
