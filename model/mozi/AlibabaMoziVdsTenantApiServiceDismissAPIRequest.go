package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziVdsTenantApiServiceDismissAPIRequest
MOZI解除组织主管服务 API请求
alibaba.mozi.vds.tenant.api.service.dismiss

解除组织主管 */
type AlibabaMoziVdsTenantApiServiceDismissAPIRequest struct {
	model.Params
	// 第一个入参
	_par0 *DismissOrganizationSupervisorRequest
}

// New
