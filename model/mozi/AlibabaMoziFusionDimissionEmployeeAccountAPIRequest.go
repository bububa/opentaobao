package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziFusionDimissionEmployeeAccountAPIRequest
人员离职 API请求
alibaba.mozi.fusion.dimission.employee.account

人员离职并且回收账号 */
type AlibabaMoziFusionDimissionEmployeeAccountAPIRequest struct {
	model.Params
	// 入参
	_dimissionEmployee *RemoveTenantEmployeeAndAccountRequest
}

// New
