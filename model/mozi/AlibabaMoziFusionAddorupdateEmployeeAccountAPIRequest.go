package mozi

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest 添加人员和账号复合接口 API请求
// alibaba.mozi.fusion.addorupdate.employee.account
//
// 添加人员和账号复合接口
type AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest struct {
	model.Params
	// 人员账号
	_employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest
}

// NewAlibabaMoziFusionAddorupdateEmployeeAccountRequest 初始化AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest对象
func NewAlibabaMoziFusionAddorupdateEmployeeAccountRequest() *AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest {
	return &AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) Reset() {
	r._employeeAccount = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) GetApiMethodName() string {
	return "alibaba.mozi.fusion.addorupdate.employee.account"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetEmployeeAccount is EmployeeAccount Setter
// 人员账号
func (r *AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) SetEmployeeAccount(_employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest) error {
	r._employeeAccount = _employeeAccount
	r.Set("employee_account", _employeeAccount)
	return nil
}

// GetEmployeeAccount EmployeeAccount Getter
func (r AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) GetEmployeeAccount() *AddOrUpdateTenantEmployeeAndAccountRequest {
	return r._employeeAccount
}

var poolAlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMoziFusionAddorupdateEmployeeAccountRequest()
	},
}

// GetAlibabaMoziFusionAddorupdateEmployeeAccountRequest 从 sync.Pool 获取 AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest
func GetAlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest() *AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest {
	return poolAlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest.Get().(*AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest)
}

// ReleaseAlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest 将 AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest 放入 sync.Pool
func ReleaseAlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest(v *AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest) {
	v.Reset()
	poolAlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest.Put(v)
}
