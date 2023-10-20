package mozivds

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceAddadminAPIRequest 新建租户管理员 API请求
// alibaba.mozi.vds.tenant.api.service.addadmin
//
// 新建租户管理员
// alibaba.mozi.vds.tenant.api.service.addadmin
type AlibabaMoziVdsTenantApiServiceAddadminAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *AddTenantAdminsRequest
}

// NewAlibabaMoziVdsTenantApiServiceAddadminRequest 初始化AlibabaMoziVdsTenantApiServiceAddadminAPIRequest对象
func NewAlibabaMoziVdsTenantApiServiceAddadminRequest() *AlibabaMoziVdsTenantApiServiceAddadminAPIRequest {
	return &AlibabaMoziVdsTenantApiServiceAddadminAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) Reset() {
	r._param0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.addadmin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam0 is Param0 Setter
// 请求参数
func (r *AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) SetParam0(_param0 *AddTenantAdminsRequest) error {
	r._param0 = _param0
	r.Set("param0", _param0)
	return nil
}

// GetParam0 Param0 Getter
func (r AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) GetParam0() *AddTenantAdminsRequest {
	return r._param0
}

var poolAlibabaMoziVdsTenantApiServiceAddadminAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziVdsTenantApiServiceAddadminRequest()
	},
}

// GetAlibabaMoziVdsTenantApiServiceAddadminRequest 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceAddadminAPIRequest
func GetAlibabaMoziVdsTenantApiServiceAddadminAPIRequest() *AlibabaMoziVdsTenantApiServiceAddadminAPIRequest {
	return poolAlibabaMoziVdsTenantApiServiceAddadminAPIRequest.Get().(*AlibabaMoziVdsTenantApiServiceAddadminAPIRequest)
}

// ReleaseAlibabaMoziVdsTenantApiServiceAddadminAPIRequest 将 AlibabaMoziVdsTenantApiServiceAddadminAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceAddadminAPIRequest(v *AlibabaMoziVdsTenantApiServiceAddadminAPIRequest) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceAddadminAPIRequest.Put(v)
}
