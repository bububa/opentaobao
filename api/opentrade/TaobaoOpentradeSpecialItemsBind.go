package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeSpecialItemsBind 专属下单场景商品绑定
// taobao.opentrade.special.items.bind
//
// 专属下单场景商品绑定
func TaobaoOpentradeSpecialItemsBind(clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialItemsBindAPIRequest, resp *opentrade.TaobaoOpentradeSpecialItemsBindAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
