package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSkuPriceUpdate 更新商品SKU的价格
// taobao.item.sku.price.update
//
// 更新商品SKU的价格
func TaobaoItemSkuPriceUpdate(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemSkuPriceUpdateAPIRequest, resp *tbitem.TaobaoItemSkuPriceUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
