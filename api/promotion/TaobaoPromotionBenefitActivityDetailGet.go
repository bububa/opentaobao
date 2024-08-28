package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivityDetailGet 活动关联的权益详情获取
// taobao.promotion.benefit.activity.detail.get
//
// 活动关联的权益详情获取
func TaobaoPromotionBenefitActivityDetailGet(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityDetailGetAPIRequest, resp *promotion.TaobaoPromotionBenefitActivityDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
