package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateDetailAPIRequest 商家协同单查询详情 API请求
// alitrip.agent.coordinate.detail
//
// 商家协同单查询详情
type AlitripAgentCoordinateDetailAPIRequest struct {
	model.Params
	// 查询协同单详情请求入参
	_coordinationDetailRequestDto *CoordinationDetailRequestDto
}

// NewAlitripAgentCoordinateDetailRequest 初始化AlitripAgentCoordinateDetailAPIRequest对象
func NewAlitripAgentCoordinateDetailRequest() *AlitripAgentCoordinateDetailAPIRequest {
	return &AlitripAgentCoordinateDetailAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentCoordinateDetailAPIRequest) Reset() {
	r._coordinationDetailRequestDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentCoordinateDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentCoordinateDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentCoordinateDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCoordinationDetailRequestDto is CoordinationDetailRequestDto Setter
// 查询协同单详情请求入参
func (r *AlitripAgentCoordinateDetailAPIRequest) SetCoordinationDetailRequestDto(_coordinationDetailRequestDto *CoordinationDetailRequestDto) error {
	r._coordinationDetailRequestDto = _coordinationDetailRequestDto
	r.Set("coordination_detail_request_dto", _coordinationDetailRequestDto)
	return nil
}

// GetCoordinationDetailRequestDto CoordinationDetailRequestDto Getter
func (r AlitripAgentCoordinateDetailAPIRequest) GetCoordinationDetailRequestDto() *CoordinationDetailRequestDto {
	return r._coordinationDetailRequestDto
}

var poolAlitripAgentCoordinateDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentCoordinateDetailRequest()
	},
}

// GetAlitripAgentCoordinateDetailRequest 从 sync.Pool 获取 AlitripAgentCoordinateDetailAPIRequest
func GetAlitripAgentCoordinateDetailAPIRequest() *AlitripAgentCoordinateDetailAPIRequest {
	return poolAlitripAgentCoordinateDetailAPIRequest.Get().(*AlitripAgentCoordinateDetailAPIRequest)
}

// ReleaseAlitripAgentCoordinateDetailAPIRequest 将 AlitripAgentCoordinateDetailAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentCoordinateDetailAPIRequest(v *AlitripAgentCoordinateDetailAPIRequest) {
	v.Reset()
	poolAlitripAgentCoordinateDetailAPIRequest.Put(v)
}
