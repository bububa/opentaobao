package mozi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziVdsTenantApiServiceTenantbyid 按租户ID查询租户信息
// alibaba.mozi.vds.tenant.api.service.tenantbyid
//
// 按租户ID查询租户信息
func AlibabaMoziVdsTenantApiServiceTenantbyid(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServiceTenantbyidAPIRequest, resp *mozi.AlibabaMoziVdsTenantApiServiceTenantbyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
