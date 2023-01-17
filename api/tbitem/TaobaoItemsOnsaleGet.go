package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemsOnsaleGet 获取当前会话用户出售中的商品列表
// taobao.items.onsale.get
//
// 获取当前用户作为卖家的出售中的商品列表，并能根据传入的搜索条件对出售中的商品列表进行过滤
// 只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get 获取
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoItemsOnsaleGet(clt *core.SDKClient, req *tbitem.TaobaoItemsOnsaleGetAPIRequest, session string) (*tbitem.TaobaoItemsOnsaleGetAPIResponse, error) {
	var resp tbitem.TaobaoItemsOnsaleGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
