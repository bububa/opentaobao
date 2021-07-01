package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest
按租户ID查询租户信息 API请求
alibaba.mozi.vds.tenant.api.service.tenantbyid

按租户ID查询租户信息 */
type AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest struct {
	model.Params
	// 入参
	_par0 *GetTenantByIdRequest
}

// New
