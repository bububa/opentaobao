package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleItemUserPublishitems 发布的商品列表
// alibaba.idle.item.user.publishitems
//
// 为服务商的卖家提供发布的闲鱼商品列表
func AlibabaIdleItemUserPublishitems(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleItemUserPublishitemsAPIRequest, resp *idleisv.AlibabaIdleItemUserPublishitemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
