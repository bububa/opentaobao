package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamozivdstenantapiservicegetadminAPIRequest 获取员工租户管理员信息（查询员工是否为租户管理员） API请求
// alibaba.mozi.vds.tenant.api.service.getadmin
//
// 获取员工租户管理员信息（查询员工是否为租户管理员）
type AlibabamozivdstenantapiservicegetadminAPIRequest struct {
	model.Params
	// 入参
	_par0 *GetEmployeeTenantAdminInfoRequest
}

// NewAlibabamozivdstenantapiservicegetadminRequest 初始化AlibabamozivdstenantapiservicegetadminAPIRequest对象
func NewAlibabamozivdstenantapiservicegetadminRequest() *AlibabamozivdstenantapiservicegetadminAPIRequest {
	return &AlibabamozivdstenantapiservicegetadminAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamozivdstenantapiservicegetadminAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.getadmin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamozivdstenantapiservicegetadminAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamozivdstenantapiservicegetadminAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 入参
func (r *AlibabamozivdstenantapiservicegetadminAPIRequest) SetPar0(_par0 *GetEmployeeTenantAdminInfoRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabamozivdstenantapiservicegetadminAPIRequest) GetPar0() *GetEmployeeTenantAdminInfoRequest {
	return r._par0
}
