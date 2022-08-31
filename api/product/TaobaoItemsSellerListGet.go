package product

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoItemsSellerListGet 批量获取商品详细信息
// taobao.items.seller.list.get
//
// 批量获取商品详细信息
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoItemsSellerListGet(clt *core.SDKClient, req *product.TaobaoItemsSellerListGetAPIRequest, session string) (*product.TaobaoItemsSellerListGetAPIResponse, error) {
	var resp product.TaobaoItemsSellerListGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
