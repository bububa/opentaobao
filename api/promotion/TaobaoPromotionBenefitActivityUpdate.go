package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionBenefitActivityUpdate 修改关联的活动权益
// taobao.promotion.benefit.activity.update
//
// 修改卖家活动中关联的对应的权益。
func TaobaoPromotionBenefitActivityUpdate(clt *core.SDKClient, req *promotion.TaobaoPromotionBenefitActivityUpdateAPIRequest, resp *promotion.TaobaoPromotionBenefitActivityUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
