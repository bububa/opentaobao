package flight

import (
	"net/url"

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
		Params: model.NewParams(),
	}
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
