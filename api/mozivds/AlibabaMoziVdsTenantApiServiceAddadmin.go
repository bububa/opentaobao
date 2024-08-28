package mozivds

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozivds"
)

// AlibabaMoziVdsTenantApiServiceAddadmin 新建租户管理员
// alibaba.mozi.vds.tenant.api.service.addadmin
//
// 新建租户管理员
// alibaba.mozi.vds.tenant.api.service.addadmin
func AlibabaMoziVdsTenantApiServiceAddadmin(ctx context.Context, clt *core.SDKClient, req *mozivds.AlibabaMoziVdsTenantApiServiceAddadminAPIRequest, resp *mozivds.AlibabaMoziVdsTenantApiServiceAddadminAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
