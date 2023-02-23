package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest 按租户ID查询租户信息 API请求
// alibaba.mozi.vds.tenant.api.service.tenantbyid
//
// 按租户ID查询租户信息
type AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest struct {
	model.Params
	// 入参
	_par0 *GetTenantByIdRequest
}

// NewAlibabaMoziVdsTenantApiServiceTenantbyidRequest 初始化AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest对象
func NewAlibabaMoziVdsTenantApiServiceTenantbyidRequest() *AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest {
	return &AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.tenantbyid"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 入参
func (r *AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest) SetPar0(_par0 *GetTenantByIdRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest) GetPar0() *GetTenantByIdRequest {
	return r._par0
}
