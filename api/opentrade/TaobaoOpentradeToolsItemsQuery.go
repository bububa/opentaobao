package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeToolsItemsQuery 交易开放获取商品绑定信息
// taobao.opentrade.tools.items.query
//
// 交易开放获取商品绑定信息
func TaobaoOpentradeToolsItemsQuery(clt *core.SDKClient, req *opentrade.TaobaoOpentradeToolsItemsQueryAPIRequest, resp *opentrade.TaobaoOpentradeToolsItemsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
