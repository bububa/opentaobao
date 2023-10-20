package tbitem

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemPublishMarketGet 获取商家可发布商品的市场信息
// alibaba.item.publish.market.get
//
// 获取商家可发布商品的市场信息
func AlibabaItemPublishMarketGet(clt *core.SDKClient, req *tbitem.AlibabaItemPublishMarketGetAPIRequest, resp *tbitem.AlibabaItemPublishMarketGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
