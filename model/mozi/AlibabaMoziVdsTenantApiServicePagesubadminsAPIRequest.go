package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozivdstenantapiservicepagesubadminsAPIRequest 分页查询租户子管理员 API请求
// alibaba.mozi.vds.tenant.api.service.pagesubadmins
//
// 分页查询租户子管理员
type AlibabamozivdstenantapiservicepagesubadminsAPIRequest struct {
	model.Params
	// 入参
	_par0 *PageTenantSubAdminsRequest
}

// NewAlibabamozivdstenantapiservicepagesubadminsRequest 初始化AlibabamozivdstenantapiservicepagesubadminsAPIRequest对象
func NewAlibabamozivdstenantapiservicepagesubadminsRequest() *AlibabamozivdstenantapiservicepagesubadminsAPIRequest {
	return &AlibabamozivdstenantapiservicepagesubadminsAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozivdstenantapiservicepagesubadminsAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.pagesubadmins"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozivdstenantapiservicepagesubadminsAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozivdstenantapiservicepagesubadminsAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 入参
func (r *AlibabamozivdstenantapiservicepagesubadminsAPIRequest) SetPar0(_par0 *PageTenantSubAdminsRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabamozivdstenantapiservicepagesubadminsAPIRequest) GetPar0() *PageTenantSubAdminsRequest {
	return r._par0
}
