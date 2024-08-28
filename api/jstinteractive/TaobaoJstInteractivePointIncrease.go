package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractivePointIncrease 互动积分发放接口
// taobao.jst.interactive.point.increase
//
// 向用户发放互动积分
func TaobaoJstInteractivePointIncrease(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractivePointIncreaseAPIRequest, resp *jstinteractive.TaobaoJstInteractivePointIncreaseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
