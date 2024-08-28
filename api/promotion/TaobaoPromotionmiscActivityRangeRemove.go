package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionmiscActivityRangeRemove 去除活动参与的商品
// taobao.promotionmisc.activity.range.remove
//
// 去除活动参与的商品
func TaobaoPromotionmiscActivityRangeRemove(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionmiscActivityRangeRemoveAPIRequest, resp *promotion.TaobaoPromotionmiscActivityRangeRemoveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
