package interact

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractIsvlotteryIsvdraw 天猫奖池鉴权接口
// alibaba.interact.isvlottery.isvdraw
//
// 鉴权接口，为tida.isvdraw接口鉴权
func AlibabaInteractIsvlotteryIsvdraw(ctx context.Context, clt *core.SDKClient, req *interact.AlibabaInteractIsvlotteryIsvdrawAPIRequest, resp *interact.AlibabaInteractIsvlotteryIsvdrawAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
