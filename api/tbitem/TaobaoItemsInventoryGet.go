package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// TaobaoItemsInventoryGet 得到当前会话用户库存中的商品列表
// taobao.items.inventory.get
//
// 获取当前用户作为卖家的仓库中的商品列表，并能根据传入的搜索条件对仓库中的商品列表进行过滤
// 只能获得商品的部分信息，商品的详细信息请通过taobao.item.seller.get获取&lt;br/&gt;
// &lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
func TaobaoItemsInventoryGet(clt *core.SDKClient, req *tbitem.TaobaoItemsInventoryGetAPIRequest, session string) (*tbitem.TaobaoItemsInventoryGetAPIResponse, error) {
	var resp tbitem.TaobaoItemsInventoryGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
