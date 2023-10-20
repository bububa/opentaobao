package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoSellercatsListGet 获取前台展示的店铺内卖家自定义商品类目
// taobao.sellercats.list.get
//
// 此API获取当前卖家店铺在淘宝前端被展示的浏览导航类目（面向买家）
func TaobaoSellercatsListGet(clt *core.SDKClient, req *shop.TaobaoSellercatsListGetAPIRequest, resp *shop.TaobaoSellercatsListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
