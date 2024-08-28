package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivityDelete 删除关联的活动权益
// taobao.promotion.benefit.activity.delete
//
// 删除关联的活动权益
func TaobaoPromotionBenefitActivityDelete(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityDeleteAPIRequest, resp *promotion.TaobaoPromotionBenefitActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
