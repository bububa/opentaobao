package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateGobackAPIRequest 协同单驳回 API请求
// alitrip.agent.coordinate.goback
//
// 协同单驳回
type AlitripAgentCoordinateGobackAPIRequest struct {
	model.Params
	// 驳回请求入参
	_gobackDto *GoBackDto
}

// NewAlitripAgentCoordinateGobackRequest 初始化AlitripAgentCoordinateGobackAPIRequest对象
func NewAlitripAgentCoordinateGobackRequest() *AlitripAgentCoordinateGobackAPIRequest {
	return &AlitripAgentCoordinateGobackAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentCoordinateGobackAPIRequest) Reset() {
	r._gobackDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentCoordinateGobackAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.goback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentCoordinateGobackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentCoordinateGobackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGobackDto is GobackDto Setter
// 驳回请求入参
func (r *AlitripAgentCoordinateGobackAPIRequest) SetGobackDto(_gobackDto *GoBackDto) error {
	r._gobackDto = _gobackDto
	r.Set("goback_dto", _gobackDto)
	return nil
}

// GetGobackDto GobackDto Getter
func (r AlitripAgentCoordinateGobackAPIRequest) GetGobackDto() *GoBackDto {
	return r._gobackDto
}

var poolAlitripAgentCoordinateGobackAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentCoordinateGobackRequest()
	},
}

// GetAlitripAgentCoordinateGobackRequest 从 sync.Pool 获取 AlitripAgentCoordinateGobackAPIRequest
func GetAlitripAgentCoordinateGobackAPIRequest() *AlitripAgentCoordinateGobackAPIRequest {
	return poolAlitripAgentCoordinateGobackAPIRequest.Get().(*AlitripAgentCoordinateGobackAPIRequest)
}

// ReleaseAlitripAgentCoordinateGobackAPIRequest 将 AlitripAgentCoordinateGobackAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentCoordinateGobackAPIRequest(v *AlitripAgentCoordinateGobackAPIRequest) {
	v.Reset()
	poolAlitripAgentCoordinateGobackAPIRequest.Put(v)
}
