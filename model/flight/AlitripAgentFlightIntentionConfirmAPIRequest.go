package flight

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripAgentFlightIntentionConfirmAPIRequest) Reset() {
	r._paramConfirmRequestDTO = nil
	r.Params.ToZero()
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

var poolAlitripAgentFlightIntentionConfirmAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripAgentFlightIntentionConfirmRequest()
	},
}

// GetAlitripAgentFlightIntentionConfirmRequest 从 sync.Pool 获取 AlitripAgentFlightIntentionConfirmAPIRequest
func GetAlitripAgentFlightIntentionConfirmAPIRequest() *AlitripAgentFlightIntentionConfirmAPIRequest {
	return poolAlitripAgentFlightIntentionConfirmAPIRequest.Get().(*AlitripAgentFlightIntentionConfirmAPIRequest)
}

// ReleaseAlitripAgentFlightIntentionConfirmAPIRequest 将 AlitripAgentFlightIntentionConfirmAPIRequest 放入 sync.Pool
func ReleaseAlitripAgentFlightIntentionConfirmAPIRequest(v *AlitripAgentFlightIntentionConfirmAPIRequest) {
	v.Reset()
	poolAlitripAgentFlightIntentionConfirmAPIRequest.Put(v)
}
