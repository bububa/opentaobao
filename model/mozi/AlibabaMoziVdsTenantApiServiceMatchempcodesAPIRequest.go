package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozivdstenantapiservicematchempcodesAPIRequest 校验组-员工是否匹配 API请求
// alibaba.mozi.vds.tenant.api.service.matchempcodes
//
// 校验组-员工是否匹配
type AlibabamozivdstenantapiservicematchempcodesAPIRequest struct {
	model.Params
	// 入参
	_par0 *MatchWithEmployeeRequest
}

// NewAlibabamozivdstenantapiservicematchempcodesRequest 初始化AlibabamozivdstenantapiservicematchempcodesAPIRequest对象
func NewAlibabamozivdstenantapiservicematchempcodesRequest() *AlibabamozivdstenantapiservicematchempcodesAPIRequest {
	return &AlibabamozivdstenantapiservicematchempcodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozivdstenantapiservicematchempcodesAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.matchempcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozivdstenantapiservicematchempcodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozivdstenantapiservicematchempcodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 入参
func (r *AlibabamozivdstenantapiservicematchempcodesAPIRequest) SetPar0(_par0 *MatchWithEmployeeRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabamozivdstenantapiservicematchempcodesAPIRequest) GetPar0() *MatchWithEmployeeRequest {
	return r._par0
}
