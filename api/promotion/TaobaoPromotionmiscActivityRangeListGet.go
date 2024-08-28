package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscActivityRangeListGet 查询活动参与的商品
// taobao.promotionmisc.activity.range.list.get
//
// 查询活动参与的商品
func TaobaoPromotionmiscActivityRangeListGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscActivityRangeListGetAPIRequest, resp *promotion.TaobaoPromotionmiscActivityRangeListGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
