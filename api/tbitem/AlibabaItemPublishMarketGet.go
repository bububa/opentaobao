package tbitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbitem"
)

// AlibabaItemPublishMarketGet 获取商家可发布商品的市场信息
// alibaba.item.publish.market.get
//
// 获取商家可发布商品的市场信息
func AlibabaItemPublishMarketGet(ctx context.Context, clt *core.SDKClient, req *tbitem.AlibabaItemPublishMarketGetAPIRequest, resp *tbitem.AlibabaItemPublishMarketGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
