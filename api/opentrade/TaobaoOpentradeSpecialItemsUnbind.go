package opentrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

// TaobaoOpentradeSpecialItemsUnbind 专属下单场景商品解绑
// taobao.opentrade.special.items.unbind
//
// 专属下单场景商品解绑
func TaobaoOpentradeSpecialItemsUnbind(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeSpecialItemsUnbindAPIRequest, resp *opentrade.TaobaoOpentradeSpecialItemsUnbindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
