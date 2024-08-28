package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemGet 根据商品ID获取商品信息
// taobao.wlb.item.get
//
// 根据商品ID获取商品信息,除了获取商品信息外还可获取商品属性信息和库存信息。
func TaobaoWlbItemGet(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbItemGetAPIRequest, resp *wlb.TaobaoWlbItemGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
