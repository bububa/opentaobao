package mozi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziVdsTenantApiServiceMatchempcodes 校验组-员工是否匹配
// alibaba.mozi.vds.tenant.api.service.matchempcodes
//
// 校验组-员工是否匹配
func AlibabaMoziVdsTenantApiServiceMatchempcodes(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziVdsTenantApiServiceMatchempcodesAPIRequest, resp *mozi.AlibabaMoziVdsTenantApiServiceMatchempcodesAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
