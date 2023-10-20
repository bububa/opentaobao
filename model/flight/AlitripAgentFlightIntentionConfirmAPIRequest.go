package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripagentflightintentionconfirmAPIRequest 意向单确认 API请求
// alitrip.agent.flight.intention.confirm
//
// 意向单确认
type AlitripagentflightintentionconfirmAPIRequest struct {
	model.Params
	// 确认入参
	_paramConfirmRequestDTO *ConfirmRequestDto
}

// NewAlitripagentflightintentionconfirmRequest 初始化AlitripagentflightintentionconfirmAPIRequest对象
func NewAlitripagentflightintentionconfirmRequest() *AlitripagentflightintentionconfirmAPIRequest {
	return &AlitripagentflightintentionconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripagentflightintentionconfirmAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.intention.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripagentflightintentionconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripagentflightintentionconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamConfirmRequestDTO is ParamConfirmRequestDTO Setter
// 确认入参
func (r *AlitripagentflightintentionconfirmAPIRequest) SetParamConfirmRequestDTO(_paramConfirmRequestDTO *ConfirmRequestDto) error {
	r._paramConfirmRequestDTO = _paramConfirmRequestDTO
	r.Set("param_confirm_request_d_t_o", _paramConfirmRequestDTO)
	return nil
}

// GetParamConfirmRequestDTO ParamConfirmRequestDTO Getter
func (r AlitripagentflightintentionconfirmAPIRequest) GetParamConfirmRequestDTO() *ConfirmRequestDto {
	return r._paramConfirmRequestDTO
}
