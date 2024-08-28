package opentrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeToolsItemsQuery 交易开放获取商品绑定信息
// taobao.opentrade.tools.items.query
//
// 交易开放获取商品绑定信息
func TaobaoOpentradeToolsItemsQuery(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeToolsItemsQueryAPIRequest, resp *opentrade.TaobaoOpentradeToolsItemsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
