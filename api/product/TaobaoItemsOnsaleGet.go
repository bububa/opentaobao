package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

/* TaobaoItemsOnsaleGet
获取当前会话用户出售中的商品列表
taobao.items.onsale.get

获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤
只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get 获取
<br/><strong><a href="https://console.open.taobao.com/dingWeb.htm?from=itemapi" target="_blank">点击查看更多商品API说明</a></strong> */
func TaobaoItemsOnsaleGet(clt *core.SDKClient, req *product.TaobaoItemsOnsaleGetAPIRequest, session string) (*product.TaobaoItemsOnsaleGetAPIResponse, error) {
	var resp product.TaobaoItemsOnsaleGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
