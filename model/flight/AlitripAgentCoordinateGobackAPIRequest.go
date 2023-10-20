package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinategobackAPIRequest 协同单驳回 API请求
// alitrip.agent.coordinate.goback
//
// 协同单驳回
type AlitripagentcoordinategobackAPIRequest struct {
	model.Params
	// 驳回请求入参
	_gobackDto *GoBackDto
}

// NewAlitripagentcoordinategobackRequest 初始化AlitripagentcoordinategobackAPIRequest对象
func NewAlitripagentcoordinategobackRequest() *AlitripagentcoordinategobackAPIRequest {
	return &AlitripagentcoordinategobackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentcoordinategobackAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.goback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentcoordinategobackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentcoordinategobackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetGobackDto is GobackDto Setter
// 驳回请求入参
func (r *AlitripagentcoordinategobackAPIRequest) SetGobackDto(_gobackDto *GoBackDto) error {
	r._gobackDto = _gobackDto
	r.Set("goback_dto", _gobackDto)
	return nil
}

// GetGobackDto GobackDto Getter
func (r AlitripagentcoordinategobackAPIRequest) GetGobackDto() *GoBackDto {
	return r._gobackDto
}
