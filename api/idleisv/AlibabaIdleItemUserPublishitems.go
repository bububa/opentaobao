package idleisv

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleItemUserPublishitems 发布的商品列表
// alibaba.idle.item.user.publishitems
//
// 为服务商的卖家提供发布的闲鱼商品列表
func AlibabaIdleItemUserPublishitems(clt *core.SDKClient, req *idleisv.AlibabaIdleItemUserPublishitemsAPIRequest, resp *idleisv.AlibabaIdleItemUserPublishitemsAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
