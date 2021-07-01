package mozi

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest
校验组-员工是否匹配 API请求
alibaba.mozi.vds.tenant.api.service.matchempcodes

校验组-员工是否匹配 */
type AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest struct {
	model.Params
	// 入参
	_par0 *MatchWithEmployeeRequest
}

// New
