package jstinteractive

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jstinteractive"
)

// TaobaoJstInteractivePointDecrease 互动积分扣减接口
// taobao.jst.interactive.point.decrease
//
// 扣减用户的互动积分
func TaobaoJstInteractivePointDecrease(ctx context.Context, clt *core.SDKClient, req *jstinteractive.TaobaoJstInteractivePointDecreaseAPIRequest, resp *jstinteractive.TaobaoJstInteractivePointDecreaseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
