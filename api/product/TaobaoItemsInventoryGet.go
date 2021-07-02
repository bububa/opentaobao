package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemsInventoryGet 得到当前会话用户库存中的商品列表
// taobao.items.inventory.get
//
// 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤
// 只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取<br/>
// <strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong>
func TaobaoItemsInventoryGet(clt *core.SDKClient, req *product.TaobaoItemsInventoryGetAPIRequest, session string) (*product.TaobaoItemsInventoryGetAPIResponse, error) {
	var resp product.TaobaoItemsInventoryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
