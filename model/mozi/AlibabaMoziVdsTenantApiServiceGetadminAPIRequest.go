package mozi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceGetadminAPIRequest 获取员工租户管理员信息（查询员工是否为租户管理员） API请求
// alibaba.mozi.vds.tenant.api.service.getadmin
//
// 获取员工租户管理员信息（查询员工是否为租户管理员）
type AlibabaMoziVdsTenantApiServiceGetadminAPIRequest struct {
	model.Params
	// 入参
	_par0 *GetEmployeeTenantAdminInfoRequest
}

// NewAlibabaMoziVdsTenantApiServiceGetadminRequest 初始化AlibabaMoziVdsTenantApiServiceGetadminAPIRequest对象
func NewAlibabaMoziVdsTenantApiServiceGetadminRequest() *AlibabaMoziVdsTenantApiServiceGetadminAPIRequest {
	return &AlibabaMoziVdsTenantApiServiceGetadminAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziVdsTenantApiServiceGetadminAPIRequest) Reset() {
	r._par0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceGetadminAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.getadmin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceGetadminAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziVdsTenantApiServiceGetadminAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 入参
func (r *AlibabaMoziVdsTenantApiServiceGetadminAPIRequest) SetPar0(_par0 *GetEmployeeTenantAdminInfoRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceGetadminAPIRequest) GetPar0() *GetEmployeeTenantAdminInfoRequest {
	return r._par0
}

var poolAlibabaMoziVdsTenantApiServiceGetadminAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziVdsTenantApiServiceGetadminRequest()
	},
}

// GetAlibabaMoziVdsTenantApiServiceGetadminRequest 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceGetadminAPIRequest
func GetAlibabaMoziVdsTenantApiServiceGetadminAPIRequest() *AlibabaMoziVdsTenantApiServiceGetadminAPIRequest {
	return poolAlibabaMoziVdsTenantApiServiceGetadminAPIRequest.Get().(*AlibabaMoziVdsTenantApiServiceGetadminAPIRequest)
}

// ReleaseAlibabaMoziVdsTenantApiServiceGetadminAPIRequest 将 AlibabaMoziVdsTenantApiServiceGetadminAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceGetadminAPIRequest(v *AlibabaMoziVdsTenantApiServiceGetadminAPIRequest) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceGetadminAPIRequest.Put(v)
}
