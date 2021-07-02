package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest 校验组-员工是否匹配 API请求
// alibaba.mozi.vds.tenant.api.service.matchempcodes
//
// 校验组-员工是否匹配
type AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest struct {
	model.Params
	// 入参
	_par0 *MatchWithEmployeeRequest
}

// NewAlibabaMoziVdsTenantApiServiceMatchempcodesRequest 初始化AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest对象
func NewAlibabaMoziVdsTenantApiServiceMatchempcodesRequest() *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest {
	return &AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.matchempcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Par0 Setter
// 入参
func (r *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) SetPar0(_par0 *MatchWithEmployeeRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// Get Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) GetPar0() *MatchWithEmployeeRequest {
	return r._par0
}
