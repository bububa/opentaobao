package mozi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionCreateEmployeeAccountAPIRequest 创建MOZI自建人员和账号 API请求
// alibaba.mozi.fusion.create.employee.account
//
// 创建MOZI自建人员和账号
type AlibabaMoziFusionCreateEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_employeeAccount *CreateTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionCreateEmployeeAccountRequest 初始化AlibabaMoziFusionCreateEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionCreateEmployeeAccountRequest() *AlibabaMoziFusionCreateEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionCreateEmployeeAccountAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziFusionCreateEmployeeAccountAPIRequest) Reset() {
	r._employeeAccount = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionCreateEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.create.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionCreateEmployeeAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziFusionCreateEmployeeAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployeeAccount is EmployeeAccount Setter
// 入参
func (r *AlibabaMoziFusionCreateEmployeeAccountAPIRequest) SetEmployeeAccount(_employeeAccount *CreateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// GetEmployeeAccount EmployeeAccount Getter
func (r AlibabaMoziFusionCreateEmployeeAccountAPIRequest) GetEmployeeAccount() *CreateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}

var poolAlibabaMoziFusionCreateEmployeeAccountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziFusionCreateEmployeeAccountRequest()
	},
}

// GetAlibabaMoziFusionCreateEmployeeAccountRequest 从 sync.Pool 获取 AlibabaMoziFusionCreateEmployeeAccountAPIRequest
func GetAlibabaMoziFusionCreateEmployeeAccountAPIRequest() *AlibabaMoziFusionCreateEmployeeAccountAPIRequest {
	return poolAlibabaMoziFusionCreateEmployeeAccountAPIRequest.Get().(*AlibabaMoziFusionCreateEmployeeAccountAPIRequest)
}

// ReleaseAlibabaMoziFusionCreateEmployeeAccountAPIRequest 将 AlibabaMoziFusionCreateEmployeeAccountAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziFusionCreateEmployeeAccountAPIRequest(v *AlibabaMoziFusionCreateEmployeeAccountAPIRequest) {
	v.Reset()
	poolAlibabaMoziFusionCreateEmployeeAccountAPIRequest.Put(v)
}
