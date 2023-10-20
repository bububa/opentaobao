package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicySpecialUploadAPIRequest 特殊政策上传 API请求
// alitrip.policy.special.upload
//
// 上传特殊类型的单程/往返政策
type AlitripPolicySpecialUploadAPIRequest struct {
	model.Params
	// 普通政策上传参数
	_paramPolicyCreateRequestDTO *PolicyCreateRequestDto
}

// NewAlitripPolicySpecialUploadRequest 初始化AlitripPolicySpecialUploadAPIRequest对象
func NewAlitripPolicySpecialUploadRequest() *AlitripPolicySpecialUploadAPIRequest {
	return &AlitripPolicySpecialUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicySpecialUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.special.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicySpecialUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPolicySpecialUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPolicyCreateRequestDTO is ParamPolicyCreateRequestDTO Setter
// 普通政策上传参数
func (r *AlitripPolicySpecialUploadAPIRequest) SetParamPolicyCreateRequestDTO(_paramPolicyCreateRequestDTO *PolicyCreateRequestDto) error {
	r._paramPolicyCreateRequestDTO = _paramPolicyCreateRequestDTO
	r.Set("param_policy_create_request_d_t_o", _paramPolicyCreateRequestDTO)
	return nil
}

// GetParamPolicyCreateRequestDTO ParamPolicyCreateRequestDTO Getter
func (r AlitripPolicySpecialUploadAPIRequest) GetParamPolicyCreateRequestDTO() *PolicyCreateRequestDto {
	return r._paramPolicyCreateRequestDTO
}
