package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionActivityGet 查询某个卖家的店铺优惠券领取活动
// taobao.promotion.activity.get
//
// 查询某个卖家的店铺优惠券领取活动&lt;br/&gt;返回，优惠券领取活动ID，优惠券ID，总领用量，每人限领量，已领取数量&lt;br/&gt;领取活动状态，优惠券领取链接&lt;br/&gt;最多50个优惠券
func TaobaoPromotionActivityGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionActivityGetAPIRequest, resp *promotion.TaobaoPromotionActivityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
