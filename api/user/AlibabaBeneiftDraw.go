package user

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/user"
)

// AlibabaBeneiftDraw 抽奖接口
// alibaba.beneift.draw
//
// 抽奖接口
func AlibabaBeneiftDraw(ctx context.Context, clt *core.SDKClient, req *user.AlibabaBeneiftDrawAPIRequest, resp *user.AlibabaBeneiftDrawAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
