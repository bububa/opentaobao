package mozi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionUpdateEmployeeAccountAPIRequest 更新人员和账号属性 API请求
// alibaba.mozi.fusion.update.employee.account
//
// 更新人员和账号基本属性
type AlibabaMoziFusionUpdateEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_employeeAccount *UpdateTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionUpdateEmployeeAccountRequest 初始化AlibabaMoziFusionUpdateEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionUpdateEmployeeAccountRequest() *AlibabaMoziFusionUpdateEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionUpdateEmployeeAccountAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) Reset() {
	r._employeeAccount = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.update.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployeeAccount is EmployeeAccount Setter
// 入参
func (r *AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) SetEmployeeAccount(_employeeAccount *UpdateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// GetEmployeeAccount EmployeeAccount Getter
func (r AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) GetEmployeeAccount() *UpdateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}

var poolAlibabaMoziFusionUpdateEmployeeAccountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziFusionUpdateEmployeeAccountRequest()
	},
}

// GetAlibabaMoziFusionUpdateEmployeeAccountRequest 从 sync.Pool 获取 AlibabaMoziFusionUpdateEmployeeAccountAPIRequest
func GetAlibabaMoziFusionUpdateEmployeeAccountAPIRequest() *AlibabaMoziFusionUpdateEmployeeAccountAPIRequest {
	return poolAlibabaMoziFusionUpdateEmployeeAccountAPIRequest.Get().(*AlibabaMoziFusionUpdateEmployeeAccountAPIRequest)
}

// ReleaseAlibabaMoziFusionUpdateEmployeeAccountAPIRequest 将 AlibabaMoziFusionUpdateEmployeeAccountAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziFusionUpdateEmployeeAccountAPIRequest(v *AlibabaMoziFusionUpdateEmployeeAccountAPIRequest) {
	v.Reset()
	poolAlibabaMoziFusionUpdateEmployeeAccountAPIRequest.Put(v)
}
