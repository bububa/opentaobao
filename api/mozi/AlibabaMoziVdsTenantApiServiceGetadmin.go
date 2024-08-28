package mozi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziVdsTenantApiServiceGetadmin 获取员工租户管理员信息（查询员工是否为租户管理员）
// alibaba.mozi.vds.tenant.api.service.getadmin
//
// 获取员工租户管理员信息（查询员工是否为租户管理员）
func AlibabaMoziVdsTenantApiServiceGetadmin(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServiceGetadminAPIRequest, resp *mozi.AlibabaMoziVdsTenantApiServiceGetadminAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
