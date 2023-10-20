package mozivds

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest 删除租户管理员服务 API请求
// alibaba.mozi.vds.tenant.api.service.removeadmin
//
// 删除租户管理员top服务
type AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest struct {
	model.Params
	// 请求入参
	_param *RemoveTenantAdminsRequest
}

// NewAlibabaMoziVdsTenantApiServiceRemoveadminRequest 初始化AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest对象
func NewAlibabaMoziVdsTenantApiServiceRemoveadminRequest() *AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest {
	return &AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest) Reset() {
	r._param = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.removeadmin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParam is Param Setter
// 请求入参
func (r *AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest) SetParam(_param *RemoveTenantAdminsRequest) error {
	r._param = _param
	r.Set("param", _param)
	return nil
}

// GetParam Param Getter
func (r AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest) GetParam() *RemoveTenantAdminsRequest {
	return r._param
}

var poolAlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziVdsTenantApiServiceRemoveadminRequest()
	},
}

// GetAlibabaMoziVdsTenantApiServiceRemoveadminRequest 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest
func GetAlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest() *AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest {
	return poolAlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest.Get().(*AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest)
}

// ReleaseAlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest 将 AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest(v *AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest.Put(v)
}
