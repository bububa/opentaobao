package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinatedetailAPIRequest 商家协同单查询详情 API请求
// alitrip.agent.coordinate.detail
//
// 商家协同单查询详情
type AlitripagentcoordinatedetailAPIRequest struct {
	model.Params
	// 查询协同单详情请求入参
	_coordinationDetailRequestDto *CoordinationDetailRequestDto
}

// NewAlitripagentcoordinatedetailRequest 初始化AlitripagentcoordinatedetailAPIRequest对象
func NewAlitripagentcoordinatedetailRequest() *AlitripagentcoordinatedetailAPIRequest {
	return &AlitripagentcoordinatedetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentcoordinatedetailAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentcoordinatedetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentcoordinatedetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCoordinationDetailRequestDto is CoordinationDetailRequestDto Setter
// 查询协同单详情请求入参
func (r *AlitripagentcoordinatedetailAPIRequest) SetCoordinationDetailRequestDto(_coordinationDetailRequestDto *CoordinationDetailRequestDto) error {
	r._coordinationDetailRequestDto = _coordinationDetailRequestDto
	r.Set("coordination_detail_request_dto", _coordinationDetailRequestDto)
	return nil
}

// GetCoordinationDetailRequestDto CoordinationDetailRequestDto Getter
func (r AlitripagentcoordinatedetailAPIRequest) GetCoordinationDetailRequestDto() *CoordinationDetailRequestDto {
	return r._coordinationDetailRequestDto
}
