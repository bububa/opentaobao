package mozi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionUpdateEmployeeAccount 更新人员和账号属性
// alibaba.mozi.fusion.update.employee.account
//
// 更新人员和账号基本属性
func AlibabaMoziFusionUpdateEmployeeAccount(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziFusionUpdateEmployeeAccountAPIRequest, resp *mozi.AlibabaMoziFusionUpdateEmployeeAccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
