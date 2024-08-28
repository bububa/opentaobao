package opentrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeToolsItemsUnbind 交易开放商品解绑
// taobao.opentrade.tools.items.unbind
//
// 交易开放商品解绑
func TaobaoOpentradeToolsItemsUnbind(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeToolsItemsUnbindAPIRequest, resp *opentrade.TaobaoOpentradeToolsItemsUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
