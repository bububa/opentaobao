package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicyruleuploadAPIRequest 规则政策上传 API请求
// alitrip.policy.rule.upload
//
// 上传特殊类型的单程/往返政策
type AlitrippolicyruleuploadAPIRequest struct {
	model.Params
	// 普通政策上传参数
	_paramPolicyCreateRequestDTO *PolicyCreateRequestDto
}

// NewAlitrippolicyruleuploadRequest 初始化AlitrippolicyruleuploadAPIRequest对象
func NewAlitrippolicyruleuploadRequest() *AlitrippolicyruleuploadAPIRequest {
	return &AlitrippolicyruleuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitrippolicyruleuploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.rule.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitrippolicyruleuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitrippolicyruleuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPolicyCreateRequestDTO is ParamPolicyCreateRequestDTO Setter
// 普通政策上传参数
func (r *AlitrippolicyruleuploadAPIRequest) SetParamPolicyCreateRequestDTO(_paramPolicyCreateRequestDTO *PolicyCreateRequestDto) error {
	r._paramPolicyCreateRequestDTO = _paramPolicyCreateRequestDTO
	r.Set("param_policy_create_request_d_t_o", _paramPolicyCreateRequestDTO)
	return nil
}

// GetParamPolicyCreateRequestDTO ParamPolicyCreateRequestDTO Getter
func (r AlitrippolicyruleuploadAPIRequest) GetParamPolicyCreateRequestDTO() *PolicyCreateRequestDto {
	return r._paramPolicyCreateRequestDTO
}
