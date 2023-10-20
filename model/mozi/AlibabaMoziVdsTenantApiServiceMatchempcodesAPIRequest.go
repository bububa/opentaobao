package mozi

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) Reset() {
	r._par0 = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.vds.tenant.api.service.matchempcodes"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPar0 is Par0 Setter
// 入参
func (r *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) SetPar0(_par0 *MatchWithEmployeeRequest) error {
	r._par0 = _par0
	r.Set("par0", _par0)
	return nil
}

// GetPar0 Par0 Getter
func (r AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) GetPar0() *MatchWithEmployeeRequest {
	return r._par0
}

var poolAlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziVdsTenantApiServiceMatchempcodesRequest()
	},
}

// GetAlibabaMoziVdsTenantApiServiceMatchempcodesRequest 从 sync.Pool 获取 AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest
func GetAlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest() *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest {
	return poolAlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest.Get().(*AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest)
}

// ReleaseAlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest 将 AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest(v *AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest) {
	v.Reset()
	poolAlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest.Put(v)
}
