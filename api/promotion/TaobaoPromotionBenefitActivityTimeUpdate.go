package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivityTimeUpdate 更新关联活动有效时间
// taobao.promotion.benefit.activity.time.update
//
// 更新关联权益的活动有效时间
func TaobaoPromotionBenefitActivityTimeUpdate(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityTimeUpdateAPIRequest, resp *promotion.TaobaoPromotionBenefitActivityTimeUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
