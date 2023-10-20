package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentcoordinaterejectAPIRequest 慧飞商家协同单拒绝接口 API请求
// alitrip.agent.coordinate.reject
//
// 慧飞商家协同单拒绝接口
type AlitripagentcoordinaterejectAPIRequest struct {
	model.Params
	// 协同单拒绝参数
	_rejectDto *RejectDto
}

// NewAlitripagentcoordinaterejectRequest 初始化AlitripagentcoordinaterejectAPIRequest对象
func NewAlitripagentcoordinaterejectRequest() *AlitripagentcoordinaterejectAPIRequest {
	return &AlitripagentcoordinaterejectAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentcoordinaterejectAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.coordinate.reject"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentcoordinaterejectAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentcoordinaterejectAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRejectDto is RejectDto Setter
// 协同单拒绝参数
func (r *AlitripagentcoordinaterejectAPIRequest) SetRejectDto(_rejectDto *RejectDto) error {
	r._rejectDto = _rejectDto
	r.Set("reject_dto", _rejectDto)
	return nil
}

// GetRejectDto RejectDto Getter
func (r AlitripagentcoordinaterejectAPIRequest) GetRejectDto() *RejectDto {
	return r._rejectDto
}
