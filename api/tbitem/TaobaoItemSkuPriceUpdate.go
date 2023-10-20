package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemSkuPriceUpdate 更新商品SKU的价格
// taobao.item.sku.price.update
//
// 更新商品SKU的价格
func TaobaoItemSkuPriceUpdate(clt *core.SDKClient, req *tbitem.TaobaoItemSkuPriceUpdateAPIRequest, session string) (*tbitem.TaobaoItemSkuPriceUpdateAPIResponse, error) {
	var resp tbitem.TaobaoItemSkuPriceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
