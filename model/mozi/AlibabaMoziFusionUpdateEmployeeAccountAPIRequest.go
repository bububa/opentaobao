package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziFusionUpdateEmployeeAccountAPIRequest
更新人员和账号属性 API请求
alibaba.mozi.fusion.update.employee.account

更新人员和账号基本属性 */
type AlibabaMoziFusionUpdateEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_employeeAccount *UpdateTenantEmployeeAndAccountRequest
}

// New
