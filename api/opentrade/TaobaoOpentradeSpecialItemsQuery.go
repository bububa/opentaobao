package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeSpecialItemsQuery 专属下单获取商品绑定信息
// taobao.opentrade.special.items.query
//
// 专属下单获取商品绑定信息
func TaobaoOpentradeSpecialItemsQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialItemsQueryAPIRequest, resp *opentrade.TaobaoOpentradeSpecialItemsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
