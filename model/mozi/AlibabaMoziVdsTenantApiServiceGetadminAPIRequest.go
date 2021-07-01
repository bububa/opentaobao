package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziVdsTenantApiServiceGetadminAPIRequest
获取员工租户管理员信息（查询员工是否为租户管理员） API请求
alibaba.mozi.vds.tenant.api.service.getadmin

获取员工租户管理员信息（查询员工是否为租户管理员） */
type AlibabaMoziVdsTenantApiServiceGetadminAPIRequest struct {
	model.Params
	// 入参
	_par0 *GetEmployeeTenantAdminInfoRequest
}

// New
