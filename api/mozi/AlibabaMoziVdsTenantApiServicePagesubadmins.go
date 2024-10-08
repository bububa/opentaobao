package mozi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziVdsTenantApiServicePagesubadmins 分页查询租户子管理员
// alibaba.mozi.vds.tenant.api.service.pagesubadmins
//
// 分页查询租户子管理员
func AlibabaMoziVdsTenantApiServicePagesubadmins(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServicePagesubadminsAPIRequest, resp *mozi.AlibabaMoziVdsTenantApiServicePagesubadminsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
