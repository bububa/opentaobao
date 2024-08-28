package mozi

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mozi"
)

// AlibabaMoziFusionCreateEmployeeAccount 创建MOZI自建人员和账号
// alibaba.mozi.fusion.create.employee.account
//
// 创建MOZI自建人员和账号
func AlibabaMoziFusionCreateEmployeeAccount(ctx context.Context, clt *core.SDKClient, req *mozi.AlibabaMoziFusionCreateEmployeeAccountAPIRequest, resp *mozi.AlibabaMoziFusionCreateEmployeeAccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
