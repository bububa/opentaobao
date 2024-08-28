package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivityUpdate 修改关联的活动权益
// taobao.promotion.benefit.activity.update
//
// 修改卖家活动中关联的对应的权益。
func TaobaoPromotionBenefitActivityUpdate(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityUpdateAPIRequest, resp *promotion.TaobaoPromotionBenefitActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
