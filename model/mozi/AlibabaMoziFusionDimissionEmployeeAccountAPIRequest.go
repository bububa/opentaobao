package mozi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionDimissionEmployeeAccountAPIRequest 人员离职 API请求
// alibaba.mozi.fusion.dimission.employee.account
//
// 人员离职并且回收账号
type AlibabaMoziFusionDimissionEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_dimissionEmployee *RemoveTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionDimissionEmployeeAccountRequest 初始化AlibabaMoziFusionDimissionEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionDimissionEmployeeAccountRequest() *AlibabaMoziFusionDimissionEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionDimissionEmployeeAccountAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) Reset() {
	r._dimissionEmployee = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.dimission.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDimissionEmployee is DimissionEmployee Setter
// 入参
func (r *AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) SetDimissionEmployee(_dimissionEmployee *RemoveTenantEmployeeAndAccountRequest) error {
	r._dimissionEmployee = _dimissionEmployee
	r.Set("dimission.employee", _dimissionEmployee)
	return nil
}

// GetDimissionEmployee DimissionEmployee Getter
func (r AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) GetDimissionEmployee() *RemoveTenantEmployeeAndAccountRequest {
	return r._dimissionEmployee
}

var poolAlibabaMoziFusionDimissionEmployeeAccountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziFusionDimissionEmployeeAccountRequest()
	},
}

// GetAlibabaMoziFusionDimissionEmployeeAccountRequest 从 sync.Pool 获取 AlibabaMoziFusionDimissionEmployeeAccountAPIRequest
func GetAlibabaMoziFusionDimissionEmployeeAccountAPIRequest() *AlibabaMoziFusionDimissionEmployeeAccountAPIRequest {
	return poolAlibabaMoziFusionDimissionEmployeeAccountAPIRequest.Get().(*AlibabaMoziFusionDimissionEmployeeAccountAPIRequest)
}

// ReleaseAlibabaMoziFusionDimissionEmployeeAccountAPIRequest 将 AlibabaMoziFusionDimissionEmployeeAccountAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziFusionDimissionEmployeeAccountAPIRequest(v *AlibabaMoziFusionDimissionEmployeeAccountAPIRequest) {
	v.Reset()
	poolAlibabaMoziFusionDimissionEmployeeAccountAPIRequest.Put(v)
}
