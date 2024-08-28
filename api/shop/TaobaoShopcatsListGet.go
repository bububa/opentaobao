package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoShopcatsListGet 获取前台展示的店铺类目
// taobao.shopcats.list.get
//
// 获取淘宝面向买家的浏览导航类目（跟后台卖家商品管理的类目有差异）
func TaobaoShopcatsListGet(ctx context.Context, clt *core.SDKClient, req *shop.TaobaoShopcatsListGetAPIRequest, resp *shop.TaobaoShopcatsListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
