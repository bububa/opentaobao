package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemPriceUpdate 更新商品价格
// taobao.item.price.update
//
// 更新商品价格
func TaobaoItemPriceUpdate(clt *core.SDKClient, req *tbitem.TaobaoItemPriceUpdateAPIRequest, session string) (*tbitem.TaobaoItemPriceUpdateAPIResponse, error) {
	var resp tbitem.TaobaoItemPriceUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
