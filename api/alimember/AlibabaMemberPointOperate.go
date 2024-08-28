package alimember

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alimember"
)

// AlibabaMemberPointOperate 积分操作
// alibaba.member.point.operate
//
// 消费会员积分
func AlibabaMemberPointOperate(ctx context.Context, clt *core.SDKClient, req *alimember.AlibabaMemberPointOperateAPIRequest, resp *alimember.AlibabaMemberPointOperateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
