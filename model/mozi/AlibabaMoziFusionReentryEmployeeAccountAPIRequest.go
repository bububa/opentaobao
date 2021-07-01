package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziFusionReentryEmployeeAccountAPIRequest
重新入职并且重新启用账号 API请求
alibaba.mozi.fusion.reentry.employee.account

重新入职并且重新启用账号 */
type AlibabaMoziFusionReentryEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_reentryEmployeeAccount *ReEntryTenantEmployeeAndAccountRequest
}

// New
