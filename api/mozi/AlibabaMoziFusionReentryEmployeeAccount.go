package mozi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionReentryEmployeeAccount 重新入职并且重新启用账号
// alibaba.mozi.fusion.reentry.employee.account
//
// 重新入职并且重新启用账号
func AlibabaMoziFusionReentryEmployeeAccount(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziFusionReentryEmployeeAccountAPIRequest, resp *mozi.AlibabaMoziFusionReentryEmployeeAccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
