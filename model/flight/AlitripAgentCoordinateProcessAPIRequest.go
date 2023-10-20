package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateProcessAPIRequest 慧飞商家协同单处理完成接口 API请求
// alitrip.agent.coordinate.process
//
// 慧飞商家协同单处理完成接口
type AlitripAgentCoordinateProcessAPIRequest struct {
	model.Params
	// 协同单处理完成参数
	_processingDto *ProcessingDto
}

// NewAlitripAgentCoordinateProcessRequest 初始化AlitripAgentCoordinateProcessAPIRequest对象
func NewAlitripAgentCoordinateProcessRequest() *AlitripAgentCoordinateProcessAPIRequest {
	return &AlitripAgentCoordinateProcessAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentCoordinateProcessAPIRequest) Reset() {
	r._processingDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentCoordinateProcessAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentCoordinateProcessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentCoordinateProcessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProcessingDto is ProcessingDto Setter
// 协同单处理完成参数
func (r *AlitripAgentCoordinateProcessAPIRequest) SetProcessingDto(_processingDto *ProcessingDto) error {
	r._processingDto = _processingDto
	r.Set("processing_dto", _processingDto)
	return nil
}

// GetProcessingDto ProcessingDto Getter
func (r AlitripAgentCoordinateProcessAPIRequest) GetProcessingDto() *ProcessingDto {
	return r._processingDto
}

var poolAlitripAgentCoordinateProcessAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentCoordinateProcessRequest()
	},
}

// GetAlitripAgentCoordinateProcessRequest 从 sync.Pool 获取 AlitripAgentCoordinateProcessAPIRequest
func GetAlitripAgentCoordinateProcessAPIRequest() *AlitripAgentCoordinateProcessAPIRequest {
	return poolAlitripAgentCoordinateProcessAPIRequest.Get().(*AlitripAgentCoordinateProcessAPIRequest)
}

// ReleaseAlitripAgentCoordinateProcessAPIRequest 将 AlitripAgentCoordinateProcessAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentCoordinateProcessAPIRequest(v *AlitripAgentCoordinateProcessAPIRequest) {
	v.Reset()
	poolAlitripAgentCoordinateProcessAPIRequest.Put(v)
}
