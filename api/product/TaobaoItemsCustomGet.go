package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TaobaoItemsCustomGet
根据外部ID取商品
taobao.items.custom.get

跟据卖家设定的商品外部id获取商品，只能获取授权卖家的商品
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
func TaobaoItemsCustomGet(clt *core.SDKClient, req *product.TaobaoItemsCustomGetAPIRequest, session string) (*product.TaobaoItemsCustomGetAPIResponse, error) {
	var resp product.TaobaoItemsCustomGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
