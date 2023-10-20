package flight

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripPolicyRuleUploadAPIRequest 规则政策上传 API请求
// alitrip.policy.rule.upload
//
// 上传特殊类型的单程/往返政策
type AlitripPolicyRuleUploadAPIRequest struct {
	model.Params
	// 普通政策上传参数
	_paramPolicyCreateRequestDTO *PolicyCreateRequestDto
}

// NewAlitripPolicyRuleUploadRequest 初始化AlitripPolicyRuleUploadAPIRequest对象
func NewAlitripPolicyRuleUploadRequest() *AlitripPolicyRuleUploadAPIRequest {
	return &AlitripPolicyRuleUploadAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripPolicyRuleUploadAPIRequest) Reset() {
	r._paramPolicyCreateRequestDTO = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripPolicyRuleUploadAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.rule.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripPolicyRuleUploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripPolicyRuleUploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPolicyCreateRequestDTO is ParamPolicyCreateRequestDTO Setter
// 普通政策上传参数
func (r *AlitripPolicyRuleUploadAPIRequest) SetParamPolicyCreateRequestDTO(_paramPolicyCreateRequestDTO *PolicyCreateRequestDto) error {
	r._paramPolicyCreateRequestDTO = _paramPolicyCreateRequestDTO
	r.Set("param_policy_create_request_d_t_o", _paramPolicyCreateRequestDTO)
	return nil
}

// GetParamPolicyCreateRequestDTO ParamPolicyCreateRequestDTO Getter
func (r AlitripPolicyRuleUploadAPIRequest) GetParamPolicyCreateRequestDTO() *PolicyCreateRequestDto {
	return r._paramPolicyCreateRequestDTO
}

var poolAlitripPolicyRuleUploadAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripPolicyRuleUploadRequest()
	},
}

// GetAlitripPolicyRuleUploadRequest 从 sync.Pool 获取 AlitripPolicyRuleUploadAPIRequest
func GetAlitripPolicyRuleUploadAPIRequest() *AlitripPolicyRuleUploadAPIRequest {
	return poolAlitripPolicyRuleUploadAPIRequest.Get().(*AlitripPolicyRuleUploadAPIRequest)
}

// ReleaseAlitripPolicyRuleUploadAPIRequest 将 AlitripPolicyRuleUploadAPIRequest 放入 sync.Pool
func ReleaseAlitripPolicyRuleUploadAPIRequest(v *AlitripPolicyRuleUploadAPIRequest) {
	v.Reset()
	poolAlitripPolicyRuleUploadAPIRequest.Put(v)
}
