package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberExit 退会
// alibaba.member.exit
//
// 商家会员解绑
func AlibabaMemberExit(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberExitAPIRequest, resp *alimember.AlibabaMemberExitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
