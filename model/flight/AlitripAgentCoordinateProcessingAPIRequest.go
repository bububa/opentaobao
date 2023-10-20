package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentCoordinateProcessingAPIRequest 慧飞商家协同单处理完成接口 API请求
// alitrip.agent.coordinate.processing
//
// 慧飞商家协同单处理完成接口
type AlitripAgentCoordinateProcessingAPIRequest struct {
	model.Params
	// 协同单处理完成参数
	_processingDto *ProcessingDto
}

// NewAlitripAgentCoordinateProcessingRequest 初始化AlitripAgentCoordinateProcessingAPIRequest对象
func NewAlitripAgentCoordinateProcessingRequest() *AlitripAgentCoordinateProcessingAPIRequest {
	return &AlitripAgentCoordinateProcessingAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentCoordinateProcessingAPIRequest) Reset() {
	r._processingDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentCoordinateProcessingAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.processing"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentCoordinateProcessingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentCoordinateProcessingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProcessingDto is ProcessingDto Setter
// 协同单处理完成参数
func (r *AlitripAgentCoordinateProcessingAPIRequest) SetProcessingDto(_processingDto *ProcessingDto) error {
	r._processingDto = _processingDto
	r.Set("processing_dto", _processingDto)
	return nil
}

// GetProcessingDto ProcessingDto Getter
func (r AlitripAgentCoordinateProcessingAPIRequest) GetProcessingDto() *ProcessingDto {
	return r._processingDto
}

var poolAlitripAgentCoordinateProcessingAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentCoordinateProcessingRequest()
	},
}

// GetAlitripAgentCoordinateProcessingRequest 从 sync.Pool 获取 AlitripAgentCoordinateProcessingAPIRequest
func GetAlitripAgentCoordinateProcessingAPIRequest() *AlitripAgentCoordinateProcessingAPIRequest {
	return poolAlitripAgentCoordinateProcessingAPIRequest.Get().(*AlitripAgentCoordinateProcessingAPIRequest)
}

// ReleaseAlitripAgentCoordinateProcessingAPIRequest 将 AlitripAgentCoordinateProcessingAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentCoordinateProcessingAPIRequest(v *AlitripAgentCoordinateProcessingAPIRequest) {
	v.Reset()
	poolAlitripAgentCoordinateProcessingAPIRequest.Put(v)
}
