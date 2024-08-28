package mozi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziVdsTenantApiServiceDismiss MOZI解除组织主管服务
// alibaba.mozi.vds.tenant.api.service.dismiss
//
// 解除组织主管
func AlibabaMoziVdsTenantApiServiceDismiss(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServiceDismissAPIRequest, resp *mozi.AlibabaMoziVdsTenantApiServiceDismissAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
