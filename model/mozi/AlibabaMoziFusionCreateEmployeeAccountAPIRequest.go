package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziFusionCreateEmployeeAccountAPIRequest
创建MOZI自建人员和账号 API请求
alibaba.mozi.fusion.create.employee.account

创建MOZI自建人员和账号 */
type AlibabaMoziFusionCreateEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_employeeAccount *CreateTenantEmployeeAndAccountRequest
}

// New
