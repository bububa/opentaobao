package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest
添加人员和账号复合接口 API请求
alibaba.mozi.fusion.addorupdate.employee.account

添加人员和账号复合接口 */
type AlibabaMoziFusionAddorupdateEmployeeAccountAPIRequest struct {
	model.Params
	// 人员账号
	_employeeAccount *AddOrUpdateTenantEmployeeAndAccountRequest
}

// New
