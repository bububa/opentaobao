package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoMiniappAdvancedTradeinfoPriceModify 高级定制商家传入改价信息
// taobao.miniapp.advanced.tradeinfo.price.modify
//
// 高级定制商家传入改价信息
func TaobaoMiniappAdvancedTradeinfoPriceModify(clt *core.SDKClient, req *opentrade.TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest, resp *opentrade.TaobaoMiniappAdvancedTradeinfoPriceModifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
