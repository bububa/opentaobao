package mozivds

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziVdsTenantApiServiceAddadminAPIRequest
新建租户管理员 API请求
alibaba.mozi.vds.tenant.api.service.addadmin

新建租户管理员
alibaba.mozi.vds.tenant.api.service.addadmin */
type AlibabaMoziVdsTenantApiServiceAddadminAPIRequest struct {
	model.Params
	// 请求参数
	_param0 *AddTenantAdminsRequest
}

// New
