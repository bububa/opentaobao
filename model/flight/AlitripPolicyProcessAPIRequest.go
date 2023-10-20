package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlitrippolicyprocessAPIRequest 政策进度查询 API请求
// alitrip.policy.process
//
// 上传特殊类型的单程/往返政策
type AlitrippolicyprocessAPIRequest struct {
	model.Params
	// 进度请求入参
	_policyTaskQueryDTO *PolicyTaskQueryDto
}

// NewAlitrippolicyprocessRequest 初始化AlitrippolicyprocessAPIRequest对象
func NewAlitrippolicyprocessRequest() *AlitrippolicyprocessAPIRequest {
	return &AlitrippolicyprocessAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitrippolicyprocessAPIRequest) GetApiMethodName() string {
	return "alitrip.policy.process"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitrippolicyprocessAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitrippolicyprocessAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPolicyTaskQueryDTO is PolicyTaskQueryDTO Setter
// 进度请求入参
func (r *AlitrippolicyprocessAPIRequest) SetPolicyTaskQueryDTO(_policyTaskQueryDTO *PolicyTaskQueryDto) error {
	r._policyTaskQueryDTO = _policyTaskQueryDTO
	r.Set("policy_task_query_d_t_o", _policyTaskQueryDTO)
	return nil
}

// GetPolicyTaskQueryDTO PolicyTaskQueryDTO Getter
func (r AlitrippolicyprocessAPIRequest) GetPolicyTaskQueryDTO() *PolicyTaskQueryDto {
	return r._policyTaskQueryDTO
}
