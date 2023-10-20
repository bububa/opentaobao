package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinateprocessAPIRequest 慧飞商家协同单处理完成接口 API请求
// alitrip.agent.coordinate.process
//
// 慧飞商家协同单处理完成接口
type AlitripagentcoordinateprocessAPIRequest struct {
	model.Params
	// 协同单处理完成参数
	_processingDto *ProcessingDto
}

// NewAlitripagentcoordinateprocessRequest 初始化AlitripagentcoordinateprocessAPIRequest对象
func NewAlitripagentcoordinateprocessRequest() *AlitripagentcoordinateprocessAPIRequest {
	return &AlitripagentcoordinateprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentcoordinateprocessAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentcoordinateprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentcoordinateprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProcessingDto is ProcessingDto Setter
// 协同单处理完成参数
func (r *AlitripagentcoordinateprocessAPIRequest) SetProcessingDto(_processingDto *ProcessingDto) error {
	r._processingDto = _processingDto
	r.Set("processing_dto", _processingDto)
	return nil
}

// GetProcessingDto ProcessingDto Getter
func (r AlitripagentcoordinateprocessAPIRequest) GetProcessingDto() *ProcessingDto {
	return r._processingDto
}
