package mozivds

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest
删除租户管理员服务 API请求
alibaba.mozi.vds.tenant.api.service.removeadmin

删除租户管理员top服务 */
type AlibabaMoziVdsTenantApiServiceRemoveadminAPIRequest struct {
	model.Params
	// 请求入参
	_param *RemoveTenantAdminsRequest
}

// New
