package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicyspecialuploadAPIRequest 特殊政策上传 API请求
// alitrip.policy.special.upload
//
// 上传特殊类型的单程/往返政策
type AlitrippolicyspecialuploadAPIRequest struct {
	model.Params
	// 普通政策上传参数
	_paramPolicyCreateRequestDTO *PolicyCreateRequestDto
}

// NewAlitrippolicyspecialuploadRequest 初始化AlitrippolicyspecialuploadAPIRequest对象
func NewAlitrippolicyspecialuploadRequest() *AlitrippolicyspecialuploadAPIRequest {
	return &AlitrippolicyspecialuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitrippolicyspecialuploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.special.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitrippolicyspecialuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitrippolicyspecialuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPolicyCreateRequestDTO is ParamPolicyCreateRequestDTO Setter
// 普通政策上传参数
func (r *AlitrippolicyspecialuploadAPIRequest) SetParamPolicyCreateRequestDTO(_paramPolicyCreateRequestDTO *PolicyCreateRequestDto) error {
	r._paramPolicyCreateRequestDTO = _paramPolicyCreateRequestDTO
	r.Set("param_policy_create_request_d_t_o", _paramPolicyCreateRequestDTO)
	return nil
}

// GetParamPolicyCreateRequestDTO ParamPolicyCreateRequestDTO Getter
func (r AlitrippolicyspecialuploadAPIRequest) GetParamPolicyCreateRequestDTO() *PolicyCreateRequestDto {
	return r._paramPolicyCreateRequestDTO
}
