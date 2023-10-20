package shop

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoSellercatsListUpdate 更新卖家自定义类目
// taobao.sellercats.list.update
//
// 此API更新卖家店铺内自定义类目 &lt;br/&gt;注：因为缓存的关系，添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
func TaobaoSellercatsListUpdate(clt *core.SDKClient, req *shop.TaobaoSellercatsListUpdateAPIRequest, resp *shop.TaobaoSellercatsListUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
