package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest 分页查询租户子管理员 API请求
// alibaba.mozi.vds.tenant.api.service.pagesubadmins
//
// 分页查询租户子管理员
type AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest struct {
	model.Params
	// 入参
	_par0 *PageTenantSubAdminsRequest
}

// NewAlibabaMoziVdsTenantApiServicePagesubadminsRequest 初始化AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest对象
func NewAlibabaMoziVdsTenantApiServicePagesubadminsRequest() *AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest {
	return &AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.pagesubadmins"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 入参
func (r *AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest) SetPar0(_par0 *PageTenantSubAdminsRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest) GetPar0() *PageTenantSubAdminsRequest {
	return r._par0
}
