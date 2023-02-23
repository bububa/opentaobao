package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyNormalUploadAPIRequest 普通政策上传 API请求
// alitrip.policy.normal.upload
//
// 上传普通类型的单程/往返政策
type AlitripPolicyNormalUploadAPIRequest struct {
	model.Params
	// 普通政策上传参数
	_paramPolicyCreateRequestDTO *PolicyCreateRequestDto
}

// NewAlitripPolicyNormalUploadRequest 初始化AlitripPolicyNormalUploadAPIRequest对象
func NewAlitripPolicyNormalUploadRequest() *AlitripPolicyNormalUploadAPIRequest {
	return &AlitripPolicyNormalUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicyNormalUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.normal.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicyNormalUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPolicyNormalUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPolicyCreateRequestDTO is ParamPolicyCreateRequestDTO Setter
// 普通政策上传参数
func (r *AlitripPolicyNormalUploadAPIRequest) SetParamPolicyCreateRequestDTO(_paramPolicyCreateRequestDTO *PolicyCreateRequestDto) error {
	r._paramPolicyCreateRequestDTO = _paramPolicyCreateRequestDTO
	r.Set("param_policy_create_request_d_t_o", _paramPolicyCreateRequestDTO)
	return nil
}

// GetParamPolicyCreateRequestDTO ParamPolicyCreateRequestDTO Getter
func (r AlitripPolicyNormalUploadAPIRequest) GetParamPolicyCreateRequestDTO() *PolicyCreateRequestDto {
	return r._paramPolicyCreateRequestDTO
}
