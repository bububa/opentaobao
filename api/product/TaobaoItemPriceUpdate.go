package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemPriceUpdate 更新商品价格
// taobao.item.price.update
//
// 更新商品价格
func TaobaoItemPriceUpdate(clt *core.SDKClient, req *product.TaobaoItemPriceUpdateAPIRequest, session string) (*product.TaobaoItemPriceUpdateAPIResponse, error) {
	var resp product.TaobaoItemPriceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
