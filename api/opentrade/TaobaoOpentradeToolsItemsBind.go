package opentrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeToolsItemsBind 交易开放商品绑定
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
func TaobaoOpentradeToolsItemsBind(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeToolsItemsBindAPIRequest, resp *opentrade.TaobaoOpentradeToolsItemsBindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
