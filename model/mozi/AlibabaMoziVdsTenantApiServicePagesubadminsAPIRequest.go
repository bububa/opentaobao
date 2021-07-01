package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest
分页查询租户子管理员 API请求
alibaba.mozi.vds.tenant.api.service.pagesubadmins

分页查询租户子管理员 */
type AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest struct {
	model.Params
	// 入参
	_par0 *PageTenantSubAdminsRequest
}

// New
