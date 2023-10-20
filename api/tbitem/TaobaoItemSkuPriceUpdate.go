package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSkuPriceUpdate 更新商品SKU的价格
// taobao.item.sku.price.update
//
// 更新商品SKU的价格
func TaobaoItemSkuPriceUpdate(clt *core.SDKClient, req *tbitem.TaobaoItemSkuPriceUpdateAPIRequest, resp *tbitem.TaobaoItemSkuPriceUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
