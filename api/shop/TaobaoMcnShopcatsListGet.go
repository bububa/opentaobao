package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoMcnShopcatsListGet 店铺类目清单
// taobao.mcn.shopcats.list.get
//
// 无需授权; 获取前台展示的店铺类目;
func TaobaoMcnShopcatsListGet(clt *core.SDKClient, req *shop.TaobaoMcnShopcatsListGetAPIRequest, resp *shop.TaobaoMcnShopcatsListGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
