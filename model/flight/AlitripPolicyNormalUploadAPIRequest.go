package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicynormaluploadAPIRequest 普通政策上传 API请求
// alitrip.policy.normal.upload
//
// 上传普通类型的单程/往返政策
type AlitrippolicynormaluploadAPIRequest struct {
	model.Params
	// 普通政策上传参数
	_paramPolicyCreateRequestDTO *PolicyCreateRequestDto
}

// NewAlitrippolicynormaluploadRequest 初始化AlitrippolicynormaluploadAPIRequest对象
func NewAlitrippolicynormaluploadRequest() *AlitrippolicynormaluploadAPIRequest {
	return &AlitrippolicynormaluploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitrippolicynormaluploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.normal.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitrippolicynormaluploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitrippolicynormaluploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPolicyCreateRequestDTO is ParamPolicyCreateRequestDTO Setter
// 普通政策上传参数
func (r *AlitrippolicynormaluploadAPIRequest) SetParamPolicyCreateRequestDTO(_paramPolicyCreateRequestDTO *PolicyCreateRequestDto) error {
	r._paramPolicyCreateRequestDTO = _paramPolicyCreateRequestDTO
	r.Set("param_policy_create_request_d_t_o", _paramPolicyCreateRequestDTO)
	return nil
}

// GetParamPolicyCreateRequestDTO ParamPolicyCreateRequestDTO Getter
func (r AlitrippolicynormaluploadAPIRequest) GetParamPolicyCreateRequestDTO() *PolicyCreateRequestDto {
	return r._paramPolicyCreateRequestDTO
}
