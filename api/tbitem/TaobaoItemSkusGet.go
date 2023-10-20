package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// Taobaoitemskusget 根据商品ID列表获取SKU信息
// taobao.item.skus.get
//
// * 获取多个商品下的所有sku
// &lt;br/&gt;&lt;strong&gt;&lt;a href=&#34;https://console.open.taobao.com/dingWeb.htm?from=itemapi&#34; target=&#34;_blank&#34;&gt;点击查看更多商品API说明&lt;/a&gt;&lt;/strong&gt;
func Taobaoitemskusget(clt *core.SDKClient, req *tbitem.TaobaoitemskusgetAPIRequest, session string) (*tbitem.TaobaoitemskusgetAPIResponse, error) {
	var resp tbitem.TaobaoitemskusgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
