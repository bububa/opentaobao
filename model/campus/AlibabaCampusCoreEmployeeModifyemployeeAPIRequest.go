package campus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaCampusCoreEmployeeModifyemployeeAPIRequest
修改员工基本信息 API请求
alibaba.campus.core.employee.modifyemployee

根据用户ID和公司ID更新员工基本信息（头像、性别、昵称） */
type AlibabaCampusCoreEmployeeModifyemployeeAPIRequest struct {
	model.Params
	// WorkBenchContext
	_workBenchContext *WorkBenchContext
	// EmployeeDto
	_employeeDto *EmployeeDto
	// 用户ID
	_accountId string
}

// New
