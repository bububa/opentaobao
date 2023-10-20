package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripAgentFlightIntentionConfirmAPIRequest 意向单确认 API请求
// alitrip.agent.flight.intention.confirm
//
// 意向单确认
type AlitripAgentFlightIntentionConfirmAPIRequest struct {
	model.Params
	// 确认入参
	_paramConfirmRequestDTO *ConfirmRequestDto
}

// NewAlitripAgentFlightIntentionConfirmRequest 初始化AlitripAgentFlightIntentionConfirmAPIRequest对象
func NewAlitripAgentFlightIntentionConfirmRequest() *AlitripAgentFlightIntentionConfirmAPIRequest {
	return &AlitripAgentFlightIntentionConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripAgentFlightIntentionConfirmAPIRequest) GetApiMethodName() string {
	return "alitrip.agent.flight.intention.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripAgentFlightIntentionConfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripAgentFlightIntentionConfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamConfirmRequestDTO is ParamConfirmRequestDTO Setter
// 确认入参
func (r *AlitripAgentFlightIntentionConfirmAPIRequest) SetParamConfirmRequestDTO(_paramConfirmRequestDTO *ConfirmRequestDto) error {
	r._paramConfirmRequestDTO = _paramConfirmRequestDTO
	r.Set("param_confirm_request_d_t_o", _paramConfirmRequestDTO)
	return nil
}

// GetParamConfirmRequestDTO ParamConfirmRequestDTO Getter
func (r AlitripAgentFlightIntentionConfirmAPIRequest) GetParamConfirmRequestDTO() *ConfirmRequestDto {
	return r._paramConfirmRequestDTO
}
