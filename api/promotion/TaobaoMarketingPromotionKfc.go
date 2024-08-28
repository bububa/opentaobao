package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoMarketingPromotionKfc 定向优惠活动名称与描述违禁词检查
// taobao.marketing.promotion.kfc
//
// 活动名称与描述违禁词检查
func TaobaoMarketingPromotionKfc(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoMarketingPromotionKfcAPIRequest, resp *promotion.TaobaoMarketingPromotionKfcAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
