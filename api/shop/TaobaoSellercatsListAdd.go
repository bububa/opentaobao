package shop

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/shop"
)

// TaobaoSellercatsListAdd 添加卖家自定义类目
// taobao.sellercats.list.add
//
// 此API添加卖家店铺内自定义类目 &lt;br/&gt;父类目parent_cid值等于0：表示此类目为店铺下的一级类目，值不等于0：表示此类目有父类目 &lt;br/&gt;注：因为缓存的关系,添加的新类目需8个小时后才可以在淘宝页面上正常显示，但是不影响在该类目下商品发布
func TaobaoSellercatsListAdd(ctx context.Context, clt *core.SDKClient, req *shop.TaobaoSellercatsListAddAPIRequest, resp *shop.TaobaoSellercatsListAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
